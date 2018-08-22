package clients

import (
	"net/http"
	"sync"
	"container/list"
	"fmt"
	"sync/atomic"
	"math"
	"io"
	"net"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"log"
)

// New creates a new RPC client based on the provided connection configuration
// details.  The notification handlers parameter may be nil if you are not
// interested in receiving notifications and will be ignored if the
// configuration is set to run in HTTP POST mode.
func NewClient(config *Config) (*Client, error) {
	// Either open a websocket connection or create an HTTP client depending
	// on the HTTP POST mode.  Also, set the notification handlers to nil
	// when running in HTTP POST mode.
	var httpClient *http.Client
	connEstablished := make(chan struct{})
	var start bool
	if config.HTTPPostMode {
		//ntfnHandlers = nil
		start = true

		var err error
		httpClient, err = newHTTPClient(config)
		if err != nil {
			return nil, err
		}
	}

	client := &Client{
		config:          config,
		httpClient:      httpClient,
		requestMap:      make(map[uint64]*list.Element),
		requestList:     list.New(),
		//ntfnHandlers:    ntfnHandlers,
		//ntfnState:       newNotificationState(),
		sendChan:        make(chan []byte, sendBufferSize),
		sendPostChan:    make(chan *sendPostDetails, sendPostBufferSize),
		connEstablished: connEstablished,
		disconnect:      make(chan struct{}),
		shutdown:        make(chan struct{}),
	}

	if start {
		log.Printf("Established connection to RPC server %s", config.Host)
		close(connEstablished)
		client.start()
	}

	return client, nil
}

// Client represents a Bitcoin RPC client which allows easy access to the
// various RPC methods available on a Bitcoin RPC server.  Each of the wrapper
// functions handle the details of converting the passed and return types to and
// from the underlying JSON types which are required for the JSON-RPC
// invocations
//
// The client provides each RPC in both synchronous (blocking) and asynchronous
// (non-blocking) forms.  The asynchronous forms are based on the concept of
// futures where they return an instance of a type that promises to deliver the
// result of the invocation at some future time.  Invoking the Receive method on
// the returned future will block until the result is available if it's not
// already.
type Client struct {
	id uint64 // atomic, so must stay 64-bit aligned

	// config holds the connection configuration assoiated with this client.
	config *Config

	// httpClient is the underlying HTTP client to use when running in HTTP
	// POST mode.
	httpClient *http.Client

	// mtx is a mutex to protect access to connection related fields.
	mtx sync.Mutex

	// disconnected indicated whether or not the server is disconnected.
	disconnected bool

	// retryCount holds the number of times the client has tried to
	// reconnect to the RPC server.
	retryCount int64

	// Track command and their response channels by ID.
	requestLock sync.Mutex
	requestMap  map[uint64]*list.Element
	requestList *list.List

	// Notifications.
	//ntfnHandlers  *NotificationHandlers
	//ntfnStateLock sync.Mutex
	//ntfnState     *notificationState

	// Networking infrastructure.
	sendChan        chan []byte
	sendPostChan    chan *sendPostDetails
	connEstablished chan struct{}
	disconnect      chan struct{}
	shutdown        chan struct{}
	wg              sync.WaitGroup
}


// NextID returns the next id to be used when sending a JSON-RPC message.  This
// ID allows responses to be associated with particular requests per the
// JSON-RPC specification.  Typically the consumer of the client does not need
// to call this function, however, if a custom request is being created and used
// this function should be used to ensure the ID is unique amongst all requests
// being made.
func (c *Client) NextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

// addRequest associates the passed jsonRequest with its id.  This allows the
// response from the remote server to be unmarshalled to the appropriate type
// and sent to the specified channel when it is received.
//
// If the client has already begun shutting down, ErrClientShutdown is returned
// and the request is not added.
//
// This function is safe for concurrent access.
func (c *Client) addRequest(jReq *jsonRequest) error {
	c.requestLock.Lock()
	defer c.requestLock.Unlock()

	// A non-blocking read of the shutdown channel with the request lock
	// held avoids adding the request to the client's internal data
	// structures if the client is in the process of shutting down (and
	// has not yet grabbed the request lock), or has finished shutdown
	// already (responding to each outstanding request with
	// ErrClientShutdown).
	select {
	case <-c.shutdown:
		return ErrClientShutdown
	default:
	}

	element := c.requestList.PushBack(jReq)
	c.requestMap[jReq.id] = element
	return nil
}

// removeRequest returns and removes the jsonRequest which contains the response
// channel and original method associated with the passed id or nil if there is
// no association.
//
// This function is safe for concurrent access.
func (c *Client) removeRequest(id uint64) *jsonRequest {
	c.requestLock.Lock()
	defer c.requestLock.Unlock()

	element := c.requestMap[id]
	if element != nil {
		delete(c.requestMap, id)
		request := c.requestList.Remove(element).(*jsonRequest)
		return request
	}

	return nil
}

// removeAllRequests removes all the jsonRequests which contain the response
// channels for outstanding requests.
//
// This function MUST be called with the request lock held.
func (c *Client) removeAllRequests() {
	c.requestMap = make(map[uint64]*list.Element)
	c.requestList.Init()
}

// trackRegisteredNtfns examines the passed command to see if it is one of
// the notification commands and updates the notification state that is used
// to automatically re-establish registered notifications on reconnects.
func (c *Client) trackRegisteredNtfns(cmd interface{}) {
	//TODO
}

// handleMessage is the main handler for incoming notifications and responses.
func (c *Client) handleMessage(msg []byte) {
	// Attempt to unmarshal the message as either a notification or response.
	var in inMessage
	in.rawResponse = new(rawResponse)
	in.rawNotification = new(rawNotification)
	err := json.Unmarshal(msg, &in)
	if err != nil {
		log.Printf("Remote server sent invalid message: %v", err)
		return
	}

	// JSON-RPC 1.0 notifications are requests with a null id.
	if in.ID == nil {
		ntfn := in.rawNotification
		if ntfn == nil {
			log.Printf("Malformed notification: missing method and parameters")
			return
		}
		if ntfn.Method == "" {
			log.Printf("Malformed notification: missing method")
			return
		}
		// params are not optional: nil isn't valid (but len == 0 is)
		if ntfn.Params == nil {
			log.Printf("Malformed notification: missing params")
			return
		}
		// Deliver the notification.
		log.Printf("Received notification [%s]", in.Method)
		//c.handleNotification(in.rawNotification)
		return
	}

	// ensure that in.ID can be converted to an integer without loss of precision
	if *in.ID < 0 || *in.ID != math.Trunc(*in.ID) {
		log.Printf("Malformed response: invalid identifier")
		return
	}

	if in.rawResponse == nil {
		log.Printf("Malformed response: missing result and error")
		return
	}

	id := uint64(*in.ID)
	log.Printf("Received response for id %d (result %s)", id, in.Result)
	request := c.removeRequest(id)

	// Nothing more to do if there is no request associated with this reply.
	if request == nil || request.responseChan == nil {
		log.Printf("Received unexpected reply: %s (id %d)", in.Result, id)
		return
	}

	// Since the command was successful, examine it to see if it's a
	// notification, and if is, add it to the notification state so it
	// can automatically be re-established on reconnect.
	c.trackRegisteredNtfns(request.cmd)

	// Deliver the response.
	result, err := in.rawResponse.result()
	request.responseChan <- &response{result: result, err: err}
}

// shouldLogReadError returns whether or not the passed error, which is expected
// to have come from reading from the websocket connection in wsInHandler,
// should be logged.
func (c *Client) shouldLogReadError(err error) bool {
	// No logging when the connetion is being forcibly disconnected.
	select {
	case <-c.shutdown:
		return false
	default:
	}

	// No logging when the connection has been disconnected.
	if err == io.EOF {
		return false
	}
	if opErr, ok := err.(*net.OpError); ok && !opErr.Temporary() {
		return false
	}

	return true
}

// wsInHandler handles all incoming messages for the websocket connection
// associated with the client.  It must be run as a goroutine.
func (c *Client) wsInHandler() {
	//TODO
}

// disconnectChan returns a copy of the current disconnect channel.  The channel
// is read protected by the client mutex, and is safe to call while the channel
// is being reassigned during a reconnect.
func (c *Client) disconnectChan() <-chan struct{} {
	c.mtx.Lock()
	ch := c.disconnect
	c.mtx.Unlock()
	return ch
}

// wsOutHandler handles all outgoing messages for the websocket connection.  It
// uses a buffered channel to serialize output messages while allowing the
// sender to continue running asynchronously.  It must be run as a goroutine.
func (c *Client) wsOutHandler() {
	//TODO
}

// sendMessage sends the passed JSON to the connected server using the
// websocket connection.  It is backed by a buffered channel, so it will not
// block until the send channel is full.
func (c *Client) sendMessage(marshalledJSON []byte) {
	// Don't send the message if disconnected.
	select {
	case c.sendChan <- marshalledJSON:
	case <-c.disconnectChan():
		return
	}
}

// reregisterNtfns creates and sends commands needed to re-establish the current
// notification state associated with the client.  It should only be called on
// on reconnect by the resendRequests function.
func (c *Client) reregisterNtfns() error {
	//TODO
	return nil
}

// resendRequests resends any requests that had not completed when the client
// disconnected.  It is intended to be called once the client has reconnected as
// a separate goroutine.
func (c *Client) resendRequests() {
	// Set the notification state back up.  If anything goes wrong,
	// disconnect the client.
	if err := c.reregisterNtfns(); err != nil {
		log.Printf("Unable to re-establish notification state: %v", err)
		c.Disconnect()
		return
	}

	// Since it's possible to block on send and more requests might be
	// added by the caller while resending, make a copy of all of the
	// requests that need to be resent now and work from the copy.  This
	// also allows the lock to be released quickly.
	c.requestLock.Lock()
	resendReqs := make([]*jsonRequest, 0, c.requestList.Len())
	var nextElem *list.Element
	for e := c.requestList.Front(); e != nil; e = nextElem {
		nextElem = e.Next()

		jReq := e.Value.(*jsonRequest)
		if _, ok := ignoreResends[jReq.method]; ok {
			// If a request is not sent on reconnect, remove it
			// from the request structures, since no reply is
			// expected.
			delete(c.requestMap, jReq.id)
			c.requestList.Remove(e)
		} else {
			resendReqs = append(resendReqs, jReq)
		}
	}
	c.requestLock.Unlock()

	for _, jReq := range resendReqs {
		// Stop resending commands if the client disconnected again
		// since the next reconnect will handle them.
		if c.Disconnected() {
			return
		}

		log.Printf("Sending command [%s] with id %d", jReq.method, jReq.id)
		c.sendMessage(jReq.marshalledJSON)
	}
}

// wsReconnectHandler listens for client disconnects and automatically tries
// to reconnect with retry interval that scales based on the number of retries.
// It also resends any commands that had not completed when the client
// disconnected so the disconnect/reconnect process is largely transparent to
// the caller.  This function is not run when the DisableAutoReconnect config
// options is set.
//
// This function must be run as a goroutine.
func (c *Client) wsReconnectHandler() {
	//TODO
}

// handleSendPostMessage handles performing the passed HTTP request, reading the
// result, unmarshalling it, and delivering the unmarshalled result to the
// provided response channel.
func (c *Client) handleSendPostMessage(details *sendPostDetails) {
	jReq := details.jsonRequest
	log.Printf("Sending command [%s] with id %d", jReq.method, jReq.id)
	httpResponse, err := c.httpClient.Do(details.httpRequest)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		jReq.responseChan <- &response{err: err}
		return
	}

	// Try to unmarshal the response as a regular JSON-RPC response.
	var resp rawResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		// When the response itself isn't a valid JSON-RPC response
		// return an error which includes the HTTP status code and raw
		// response bytes.
		err = fmt.Errorf("status code: %d, response: %q", httpResponse.StatusCode, string(respBytes))
		jReq.responseChan <- &response{err: err}
		return
	}

	res, err := resp.result()
	jReq.responseChan <- &response{result: res, err: err}
}

// sendPostHandler handles all outgoing messages when the client is running
// in HTTP POST mode.  It uses a buffered channel to serialize output messages
// while allowing the sender to continue running asynchronously.  It must be run
// as a goroutine.
func (c *Client) sendPostHandler() {
out:
	for {
		// Send any messages ready for send until the shutdown channel is closed.
		select {
		case details := <-c.sendPostChan:
			c.handleSendPostMessage(details)

		case <-c.shutdown:
			break out
		}
	}

	// Drain any wait channels before exiting so nothing is left waiting
	// around to send.
cleanup:
	for {
		select {
		case details := <-c.sendPostChan:
			details.jsonRequest.responseChan <- &response{
				result: nil,
				err:    ErrClientShutdown,
			}

		default:
			break cleanup
		}
	}
	c.wg.Done()
	log.Printf("RPC client send handler done for %s", c.config.Host)

}

// sendPostRequest sends the passed HTTP request to the RPC server using the
// HTTP client associated with the client.  It is backed by a buffered channel,
// so it will not block until the send channel is full.
func (c *Client) sendPostRequest(httpReq *http.Request, jReq *jsonRequest) {
	// Don't send the message if shutting down.
	select {
	case <-c.shutdown:
		jReq.responseChan <- &response{result: nil, err: ErrClientShutdown}
	default:
	}

	c.sendPostChan <- &sendPostDetails{
		jsonRequest: jReq,
		httpRequest: httpReq,
	}
}

// sendPost sends the passed request to the server by issuing an HTTP POST
// request using the provided response channel for the reply.  Typically a new
// connection is opened and closed for each command when using this method,
// however, the underlying HTTP client might coalesce multiple commands
// depending on several factors including the remote server configuration.
func (c *Client) sendPost(jReq *jsonRequest) {
	// Generate a request to the configured RPC server.
	protocol := "http"
	if !c.config.DisableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host
	bodyReader := bytes.NewReader(jReq.marshalledJSON)
	httpReq, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		jReq.responseChan <- &response{result: nil, err: err}
		return
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")

	// Configure basic access authorization.
	httpReq.SetBasicAuth(c.config.User, c.config.Pass)

	log.Printf("Sending command [%s] with id %d", jReq.method, jReq.id)
	c.sendPostRequest(httpReq, jReq)
}

// sendRequest sends the passed json request to the associated server using the
// provided response channel for the reply.  It handles both websocket and HTTP
// POST mode depending on the configuration of the client.
func (c *Client) sendRequest(jReq *jsonRequest) {
	// Choose which marshal and send function to use depending on whether
	// the client running in HTTP POST mode or not.  When running in HTTP
	// POST mode, the command is issued via an HTTP client.  Otherwise,
	// the command is issued via the asynchronous websocket channels.
	if c.config.HTTPPostMode {
		c.sendPost(jReq)
		return
	}

	// Check whether the websocket connection has never been established,
	// in which case the handler goroutines are not running.
	select {
	case <-c.connEstablished:
	default:
		jReq.responseChan <- &response{err: ErrClientNotConnected}
		return
	}

	// Add the request to the internal tracking map so the response from the
	// remote server can be properly detected and routed to the response
	// channel.  Then send the marshalled request via the websocket
	// connection.
	if err := c.addRequest(jReq); err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}
	log.Printf("Sending command [%s] with id %d", jReq.method, jReq.id)
	c.sendMessage(jReq.marshalledJSON)
}

// sendCmd sends the passed command to the associated server and returns a
// response channel on which the reply will be delivered at some point in the
// future.  It handles both websocket and HTTP POST mode depending on the
// configuration of the client.
func (c *Client) sendCmd(cmd interface{}) chan *response {
	// Get the method associated with the command.
	method, err := CmdMethod(cmd)
	if err != nil {
		return newFutureError(err)
	}

	// Marshal the command.
	id := c.NextID()
	marshalledJSON, err := MarshalCmd(id, cmd)
	if err != nil {
		return newFutureError(err)
	}

	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		id:             id,
		method:         method,
		cmd:            cmd,
		marshalledJSON: marshalledJSON,
		responseChan:   responseChan,
	}
	c.sendRequest(jReq)

	return responseChan
}

// sendCmdAndWait sends the passed command to the associated server, waits
// for the reply, and returns the result from it.  It will return the error
// field in the reply if there is one.
func (c *Client) sendCmdAndWait(cmd interface{}) (interface{}, error) {
	// Marshal the command to JSON-RPC, send it to the connected server, and
	// wait for a response on the returned channel.
	return receiveFuture(c.sendCmd(cmd))
}

// Disconnected returns whether or not the server is disconnected.  If a
// websocket client was created but never connected, this also returns false.
func (c *Client) Disconnected() bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	select {
	case <-c.connEstablished:
		return c.disconnected
	default:
		return false
	}
}

// doDisconnect disconnects the websocket associated with the client if it
// hasn't already been disconnected.  It will return false if the disconnect is
// not needed or the client is running in HTTP POST mode.
//
// This function is safe for concurrent access.
func (c *Client) doDisconnect() bool {
	if c.config.HTTPPostMode {
		return false
	}

	c.mtx.Lock()
	defer c.mtx.Unlock()

	// Nothing to do if already disconnected.
	if c.disconnected {
		return false
	}

	log.Printf("Disconnecting RPC client %s \n", c.config.Host)
	close(c.disconnect)
	//if c.wsConn != nil {
	//	c.wsConn.Close()
	//}
	c.disconnected = true
	return true
}

// doShutdown closes the shutdown channel and logs the shutdown unless shutdown
// is already in progress.  It will return false if the shutdown is not needed.
//
// This function is safe for concurrent access.
func (c *Client) doShutdown() bool {
	// Ignore the shutdown request if the client is already in the process
	// of shutting down or already shutdown.
	select {
	case <-c.shutdown:
		return false
	default:
	}

	log.Printf("Shutting down RPC client %s \n", c.config.Host)
	close(c.shutdown)
	return true
}

// Disconnect disconnects the current websocket associated with the client.  The
// connection will automatically be re-established unless the client was
// created with the DisableAutoReconnect flag.
//
// This function has no effect when the client is running in HTTP POST mode.
func (c *Client) Disconnect() {
	// Nothing to do if already disconnected or running in HTTP POST mode.
	if !c.doDisconnect() {
		return
	}

	c.requestLock.Lock()
	defer c.requestLock.Unlock()

	// When operating without auto reconnect, send errors to any pending
	// requests and shutdown the client.
	if c.config.DisableAutoReconnect {
		for e := c.requestList.Front(); e != nil; e = e.Next() {
			req := e.Value.(*jsonRequest)
			req.responseChan <- &response{
				result: nil,
				err:    ErrClientDisconnect,
			}
		}
		c.removeAllRequests()
		c.doShutdown()
	}
}

// Shutdown shuts down the client by disconnecting any connections associated
// with the client and, when automatic reconnect is enabled, preventing future
// attempts to reconnect.  It also stops all goroutines.
func (c *Client) Shutdown() {
	// Do the shutdown under the request lock to prevent clients from
	// adding new requests while the client shutdown process is initiated.
	c.requestLock.Lock()
	defer c.requestLock.Unlock()

	// Ignore the shutdown request if the client is already in the process
	// of shutting down or already shutdown.
	if !c.doShutdown() {
		return
	}

	// Send the ErrClientShutdown error to any pending requests.
	for e := c.requestList.Front(); e != nil; e = e.Next() {
		req := e.Value.(*jsonRequest)
		req.responseChan <- &response{
			result: nil,
			err:    ErrClientShutdown,
		}
	}
	c.removeAllRequests()

	// Disconnect the client if needed.
	c.doDisconnect()
}

// start begins processing input and output messages.
func (c *Client) start() {
	log.Printf("Starting RPC client %s", c.config.Host)

	// Start the I/O processing handlers depending on whether the client is
	// in HTTP POST mode or the default websocket mode.
	if c.config.HTTPPostMode {
		c.wg.Add(1)
		go c.sendPostHandler()
	} else {
		//c.wg.Add(3)
		//go func() {
		//	if c.ntfnHandlers != nil {
		//		if c.ntfnHandlers.OnClientConnected != nil {
		//			c.ntfnHandlers.OnClientConnected()
		//		}
		//	}
		//	c.wg.Done()
		//}()
		//go c.wsInHandler()
		//go c.wsOutHandler()
	}
}

// WaitForShutdown blocks until the client goroutines are stopped and the
// connection is closed.
func (c *Client) WaitForShutdown() {
	c.wg.Wait()
}

// Connect establishes the initial websocket connection.  This is necessary when
// a client was created after setting the DisableConnectOnNew field of the
// Config struct.
//
// Up to tries number of connections (each after an increasing backoff) will
// be tried if the connection can not be established.  The special value of 0
// indicates an unlimited number of connection attempts.
//
// This method will error if the client is not configured for websockets, if the
// connection has already been established, or if none of the connection
// attempts were successful.
func (c *Client) Connect(tries int) error {
	//TODO
	return nil
}
