package clients

import (
	"time"
	"net/http"
	"encoding/json"
)

const (
	// sendBufferSize is the number of elements the websocket send channel
	// can queue before blocking.
	sendBufferSize = 50

	// sendPostBufferSize is the number of elements the HTTP POST send
	// channel can queue before blocking.
	sendPostBufferSize = 100

	// connectionRetryInterval is the amount of time to wait in between
	// retries when automatically reconnecting to an RPC server.
	connectionRetryInterval = time.Second * 5
)

// ignoreResends is a set of all methods for requests that are "long running"
// are not be reissued by the client on reconnect.
var ignoreResends = map[string]struct{}{
	"rescan": {},
}

// sendPostDetails houses an HTTP POST request to send to an RPC server as well
// as the original JSON-RPC command and a channel to reply on when the server
// responds with the result.
type sendPostDetails struct {
	httpRequest *http.Request
	jsonRequest *jsonRequest
}

// jsonRequest holds information about a json request that is used to properly
// detect, interpret, and deliver a reply to it.
type jsonRequest struct {
	id             uint64
	method         string
	cmd            interface{}
	marshalledJSON []byte
	responseChan   chan *response
}

// response is the raw bytes of a JSON-RPC result, or the error if the response
// error object was non-null.
type response struct {
	result []byte
	err    error
}

type (
	// inMessage is the first type that an incoming message is unmarshaled
	// into. It supports both requests (for notification support) and
	// responses.  The partially-unmarshaled message is a notification if
	// the embedded ID (from the response) is nil.  Otherwise, it is a
	// response.
	inMessage struct {
		ID *float64 `json:"id"`
		*rawNotification
		*rawResponse
	}

	// rawNotification is a partially-unmarshaled JSON-RPC notification.
	rawNotification struct {
		Method string            `json:"method"`
		Params []json.RawMessage `json:"params"`
	}

	// rawResponse is a partially-unmarshaled JSON-RPC response.  For this
	// to be valid (according to JSON-RPC 1.0 spec), ID may not be nil.
	rawResponse struct {
		Result json.RawMessage   `json:"result"`
		Error  *RPCError `json:"error"`
	}
)

// result checks whether the unmarshaled response contains a non-nil error,
// returning an unmarshaled btcjson.RPCError (or an unmarshaling error) if so.
// If the response is not an error, the raw bytes of the request are
// returned for further unmashaling into specific result types.
func (r rawResponse) result() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}


// receiveFuture receives from the passed futureResult channel to extract a
// reply or any errors.  The examined errors include an error in the
// futureResult and the error in the reply from the server.  This will block
// until the result is available on the passed channel.
func receiveFuture(f chan *response) ([]byte, error) {
	// Wait for a response on the returned channel.
	r := <-f
	return r.result, r.err
}

// newFutureError returns a new future result channel that already has the
// passed error waitin on the channel with the reply set to nil.  This is useful
// to easily return errors from the various Async functions.
func newFutureError(err error) chan *response {
	responseChan := make(chan *response, 1)
	responseChan <- &response{err: err}
	return responseChan
}




// Bool is a helper routine that allocates a new bool value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// Uint is a helper routine that allocates a new uint value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Uint(v uint) *uint {
	p := new(uint)
	*p = v
	return p
}

// Int32 is a helper routine that allocates a new int32 value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Int32(v int32) *int32 {
	p := new(int32)
	*p = v
	return p
}

// Uint32 is a helper routine that allocates a new uint32 value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Uint32(v uint32) *uint32 {
	p := new(uint32)
	*p = v
	return p
}

// Int64 is a helper routine that allocates a new int64 value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Int64(v int64) *int64 {
	p := new(int64)
	*p = v
	return p
}

// Uint64 is a helper routine that allocates a new uint64 value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Uint64(v uint64) *uint64 {
	p := new(uint64)
	*p = v
	return p
}

// Float64 is a helper routine that allocates a new float64 value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func Float64(v float64) *float64 {
	p := new(float64)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value to store v and
// returns a pointer to it.  This is useful when assigning optional parameters.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}