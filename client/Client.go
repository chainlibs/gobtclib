package client

import (
	"net/http"
	"sync"
	"sync/atomic"
	glog "log"
	"io/ioutil"
	"fmt"
	"bytes"
	"github.com/chainlibs/gobtclib/utils"
	"encoding/json"
	"github.com/gobasis/log"
)

/*
Description:
New creates a new RPC client based on the provided connection configuration details.
 * Author: architect.bian
 * Date: 2018/08/23 14:53
 */
func NewClient(config *Config) (*Client) {
	var httpClient *http.Client
	var err error
	httpClient, err = utils.Http.NewClient(config.Proxy, config.EnableTLS, config.Certificates)
	if err != nil {
		glog.Fatal(err)
		return nil
	}

	client := &Client{
		config:          config,
		httpClient:      httpClient,
		requestsChan:    make(chan *requestDetail, requestsChanBufferSize),
		disconnect:      make(chan struct{}),
		shutdown:        make(chan struct{}),
	}

	return client
}

/*
Description:
Client represents a Bitcoin RPC client which allows easy access to the various RPC methods available
on a Bitcoin RPC server. Each of the wrapper functions handle the details of converting the passed and
return types th and from the underlying JSON types which are required for the JSON-RPC invocations.

The client provides each RPC in both synchronous (blocking) and asynchronous (non-blocking) forms.
The asynchronous forms are based on the concept of futures where they return an instance of a type that
promises to deliver the result of the invocation at some future time. Invoking the Receive method on
the returned future will block util the result is available if it's not already.
 * Author: architect.bian
 * Date: 2018/08/26 14:03
 */
type Client struct {
	id uint64 // atomic, so must stay 64-bit aligned
	// config holds the connection configuration assoiated with this client.
	config *Config
	// httpClient is the underlying HTTP client to use when running in HTTP POST mode.
	httpClient *http.Client
	// mtx is a mutex to protect access to connection related fields.
	mtx sync.Mutex
	// disconnected indicated whether or not the server is disconnected.
	disconnected bool
	// Networking infrastructure.
	requestsChan chan *requestDetail
	disconnect      chan struct{}
	shutdown        chan struct{}
	wg              sync.WaitGroup
}

/*
Description:
start begins processing input and output messages.
 * Author: architect.bian
 * Date: 2018/08/23 16:02
 */
func (c *Client) Startup() *Client {
	log.Info("Starting RPC client", "Host", c.config.Host)
	c.wg.Add(1)
	go c.sendPostHandler()
	return c
}

/*
Description:
sendPostHandler handles all outgoing messages when the client is running
in HTTP POST mode.  It uses a buffered channel to serialize output messages
while allowing the sender to continue running asynchronously.  It must be run
as a goroutine.
 * Author: architect.bian
 * Date: 2018/08/26 16:05
 */
func (c *Client) sendPostHandler() {
out:
	for {
		// Send any messages ready for send until the shutdown channel is closed.
		select {
		case req := <-c.requestsChan:
			c.handleRequestDetail(req)

		case <-c.shutdown:
			break out
		}
	}

	// Drain any wait channels before exiting so nothing is left waiting
	// around to send.
cleanup:
	for {
		select {
		case detail := <-c.requestsChan:
			detail.jsonRequest.responseChan <- &response{
				result: nil,
				err:    ErrClientShutdown,
			}
		default:
			break cleanup
		}
	}
	c.wg.Done()
	log.Info("RPC client send task clean up", "host", c.config.Host)
}

/*
Description:
handleRequestDetail handles performing the passed HTTP request, reading the
result, unmarshalling it, and delivering the unmarshalled result to the
provided response channel.
 * Author: architect.bian
 * Date: 2018/10/07 18:26
 */
func (c *Client) handleRequestDetail(request *requestDetail) {
	jsonReq := request.jsonRequest
	log.Debug("Sending command", "command", jsonReq.method, "id", jsonReq.id)
	httpResponse, err := c.httpClient.Do(request.httpRequest)
	if err != nil {
		jsonReq.responseChan <- &response{err: err}
		return
	}
	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		jsonReq.responseChan <- &response{err: err}
		return
	}
	log.Debug("receive result data after posted", "result", string(respBytes))
	// Try to unmarshal the response as a regular JSON-RPC response.
	var resp rawResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		// When the response itself isn't a valid JSON-RPC response
		// return an error which includes the HTTP status code and raw
		// response bytes.
		err = fmt.Errorf("status code: %d, response: %q", httpResponse.StatusCode, string(respBytes))
		jsonReq.responseChan <- &response{err: err}
		return
	}
	res, err := resp.result()
	jsonReq.responseChan <- &response{result: res, err: err}
}

/*
Description:
sendPostRequest sends the passed HTTP request to the RPC server using the
HTTP client associated with the client.  It is backed by a buffered channel,
so it will not block until the send channel is full.
 * Author: architect.bian
 * Date: 2018/08/26 19:29
 */
func (c *Client) sendPostRequest(req *http.Request, jsonReq *jsonRequest) {
	// Don't send the request if shutting down.
	select {
	case <-c.shutdown:
		jsonReq.responseChan <- &response{result: nil, err: ErrClientShutdown}
	default:
	}

	c.requestsChan <- &requestDetail{
		jsonRequest: jsonReq,
		httpRequest: req,
	}
}

/*
Description:
sendPost sends the passed request to the server by issuing an HTTP POST
request using the provided response channel for the reply.  Typically a new
connection is opened and closed for each command when using this method,
however, the underlying HTTP client might coalesce multiple commands
depending on several factors including the remote server configuration.
 * Author: architect.bian
 * Date: 2018/08/26 19:29
 */
func (c *Client) sendPost(jsonReq *jsonRequest) {
	// Generate a request to the configured RPC server.
	protocol := "http"
	if c.config.EnableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host
	bodyReader := bytes.NewReader(jsonReq.marshalledJSON)
	request, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		jsonReq.responseChan <- &response{result: nil, err: err}
		return
	}
	request.Close = true
	request.Header.Set("Content-Type", "application/json")
	// Configure basic access authorization.
	request.SetBasicAuth(c.config.User, c.config.Pass)
	c.sendPostRequest(request, jsonReq)
}

/*
Description:
sendRequest sends the passed json request to the associated server using the
provided response channel for the reply.
 * Author: architect.bian
 * Date: 2018/08/26 19:19
 */
func (c *Client) sendRequest(jReq *jsonRequest) {
	// Choose which marshal and send function to use depending on whether
	// the client running in HTTP POST mode or not.  When running in HTTP
	// POST mode, the command is issued via an HTTP client.  Otherwise,
	// the command is issued via the asynchronous websocket channels.
	c.sendPost(jReq)
	// ws ignore ...
}

/*
Description:
NextID returns the next id to be used when sending a JSON-RPC message.  This
ID allows responses to be associated with particular requestsChan per the
JSON-RPC specification.  Typically the consumer of the client does not need
to call this function, however, if a custom request is being created and used
this function should be used to ensure the ID is unique amongst all requestsChan
being made.
 * Author: architect.bian
 * Date: 2018/08/26 14:30
 */
func (c *Client) NextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

/*
Description:
sendCmd sends the passed command to the associated server and returns a
response channel on which the reply will be delivered at some point in the
future.  It handles both websocket and HTTP POST mode depending on the
configuration of the client.
 * Author: architect.bian
 * Date: 2018/08/26 18:46
 */
func (c *Client) sendCmd(cmd *Command) chan *response {
	// Marshal the command.
	id := c.NextID()
	marshalledJSON, err := MarshalCmd(id, cmd)
	if err != nil {
		return newFutureError(err)
	}
	log.Debug("posting json content", "json", string(marshalledJSON))
	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		id:             id,
		method:         cmd.name,
		cmd:            cmd,
		marshalledJSON: marshalledJSON,
		responseChan:   responseChan,
	}
	c.sendRequest(jReq)

	return responseChan
}

/*
Description:
sendCmdAndWait sends the passed command to the associated server, waits
for the reply, and returns the result from it.  It will return the error
field in the reply if there is one.
 * Author: architect.bian
 * Date: 2018/10/06 19:50
 */
func (c *Client) sendCmdAndWait(cmd *Command) (interface{}, error) {
	// Marshal the command to JSON-RPC, send it to the connected server, and
	// wait for a response on the returned channel.
	return receiveFuture(c.sendCmd(cmd))
}

/*
Description:
Shutdown closes the shutdown channel and logs the shutdown unless shutdown
is already in progress.  It will return false if the shutdown is not needed.
This function is safe for concurrent access.
 * Author: architect.bian
 * Date: 2018/08/26 19:38
 */
func (c *Client) Shutdown() {
	// Ignore the shutdown request if the client is already in the process
	// of shutting down or already shutdown.
	select {
	case <-c.shutdown:
		return
	default:
	}

	close(c.shutdown)
	c.wg.Wait() //blocks until the client goroutines are stopped and the connection is closed.
	log.Info("Shut down RPC client", "host", c.config.Host)
}
