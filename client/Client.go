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
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/futures"
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
		requestsChan:    make(chan *base.RequestDetail, requestsChanBufferSize),
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
	requestsChan chan *base.RequestDetail
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
	go c.handleRequests()
	return c
}

/*
Description:
SendAsync send command and args, return an instance of a future type
 * Author: architect.bian
 * Date: 2018/08/23 16:02
 */
func (c *Client) SendAsync(command string, args ...interface{}) futures.FutureResult {
	cmd := NewCommand(command, args...)
	return futures.FutureResult(c.sendCmd(cmd))
}

/*
Description:
Send send any command and arguments, then return a result of interface type
 * Author: architect.bian
 * Date: 2018/08/23 16:02
 */
func (c *Client) Send(command string, args ...interface{}) (*interface{}, error) {
	cmd := NewCommand(command, args...)
	return futures.FutureResult(c.sendCmd(cmd)).Receive()
}

/*
Description:
handleRequests handles all outgoing messages when the client is running
in HTTP POST mode.  It uses a buffered channel to serialize output messages
while allowing the sender to continue running asynchronously.  It must be run
as a goroutine.
 * Author: architect.bian
 * Date: 2018/08/26 16:05
 */
func (c *Client) handleRequests() {
out:
	for {
		// Send any messages ready for send until the shutdown channel is closed.
		select {
		case req := <-c.requestsChan:
			c.doRequest(req)

		case <-c.shutdown:
			break out
		}
	}

	// Drain any wait channels before exiting so nothing is left waiting
	// around to send.
cleanup:
	for {
		select {
		case req := <-c.requestsChan:
			req.Detail.ResponseChan <- &base.Response{
				Result: nil,
				Err:    ErrClientShutdown,
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
handlebase.RequestDetail handles performing the passed HTTP request, reading the
result, unmarshalling it, and delivering the unmarshalled result to the
provided response channel.
 * Author: architect.bian
 * Date: 2018/10/07 18:26
 */
func (c *Client) doRequest(requestDetail *base.RequestDetail) {
	detail := requestDetail.Detail
	log.Debug("Sending command", "command", detail.Method, "id", detail.Id)
	response, err := c.httpClient.Do(requestDetail.HttpRequest)
	if err != nil {
		detail.ResponseChan <- &base.Response{Err: err}
		return
	}
	// Read the raw bytes and close the response.
	bodyBytes, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		detail.ResponseChan <- &base.Response{Err: err}
		return
	}
	log.Debug("receive result data after posted", "result", string(bodyBytes))
	// Try to unmarshal the response as a regular JSON-RPC response.
	var respRaw base.RespRaw
	err = json.Unmarshal(bodyBytes, &respRaw)
	if err != nil {
		// When the response itself isn't a valid JSON-RPC response
		// return an error which includes the HTTP status code and raw
		// response bytes.
		err = fmt.Errorf("status code: %d, response: %q", response.StatusCode, string(bodyBytes))
		detail.ResponseChan <- &base.Response{Err: err}
		return
	}
	respBytes, err := respRaw.GetRaw()
	detail.ResponseChan <- &base.Response{Result: respBytes, Err: err}
}

/*
Description:
sendRequest sends the passed HTTP request to the RPC server using the
HTTP client associated with the client.  It is backed by a buffered channel,
so it will not block until the send channel is full.
 * Author: architect.bian
 * Date: 2018/08/26 19:29
 */
func (c *Client) sendRequest(req *http.Request, detail *base.JsonDetail) {
	// Don't send the request if shutting down.
	select {
	case <-c.shutdown:
		detail.ResponseChan <- &base.Response{Result: nil, Err: ErrClientShutdown}
	default:
	}

	c.requestsChan <- &base.RequestDetail{
		HttpRequest: req,
		Detail:      detail,
	}
}

/*
Description:
sendDetail sends the passed request to the server by issuing an HTTP POST
request using the provided response channel for the reply.  Typically a new
connection is opened and closed for each command when using this method,
however, the underlying HTTP client might coalesce multiple commands
depending on several factors including the remote server configuration.
 * Author: architect.bian
 * Date: 2018/08/26 19:29
 */
func (c *Client) sendDetail(detail *base.JsonDetail) {
	// Generate a request to the configured RPC server.
	protocol := "http"
	if c.config.EnableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host
	bodyReader := bytes.NewReader(detail.MarshalledJSON)
	request, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		detail.ResponseChan <- &base.Response{Result: nil, Err: err}
		return
	}
	request.Close = true
	request.Header.Set("Content-Type", "application/json")
	// Configure basic access authorization.
	request.SetBasicAuth(c.config.User, c.config.Pass)
	c.sendRequest(request, detail)
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
func (c *Client) sendCmd(cmd *Command) chan *base.Response {
	// Marshal the command.
	id := c.NextID()
	marshalledJSON, err := MarshalCmdToJRPC(id, cmd)
	if err != nil {
		return futures.NewFutureError(err)
	}
	log.Debug("posting json content", "json", string(marshalledJSON))
	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *base.Response, 1)
	detail := &base.JsonDetail{
		Id:             id,
		Method:         cmd.name,
		Cmd:            cmd,
		MarshalledJSON: marshalledJSON,
		ResponseChan:   responseChan,
	}
	c.sendDetail(detail)

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
	return futures.ReceiveFuture(c.sendCmd(cmd))
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
