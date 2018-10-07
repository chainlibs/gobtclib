package base

import (
	"net/http"
	"encoding/json"
)

/*
Description:
requestDetail houses an HTTP POST request to send to an RPC server as well
as the original JSON-RPC command and a channel to reply on when the server
responds with the result.
 * Author: architect.bian
 * Date: 2018/08/26 15:42
 */
type RequestDetail struct {
	HttpRequest *http.Request
	JsonRequest *JsonRequest
}

/*
Description:
jsonRequest holds information about a json request that is used to properly
detect, interpret, and deliver a reply to it.
 * Author: architect.bian
 * Date: 2018/08/26 15:43
 */
type JsonRequest struct { //TODO JsonDetail?JsonRequestDetail?
	Id             uint64
	Method         string
	Cmd            interface{}
	MarshalledJSON []byte
	ResponseChan   chan *Response // Generate the request and send it along with a channel to respond on.
}

/*
Description:
response is the raw bytes of a JSON-RPC result, or the error if the response
error object was non-null.
 * Author: architect.bian
 * Date: 2018/08/26 15:44
 */
type Response struct {
	Result []byte
	Err    error
}


/*
Description:
rawResponse is a partially-unmarshaled JSON-RPC response.  For this
to be valid (according to JSON-RPC 1.0 spec), ID may not be nil.
 * Author: architect.bian
 * Date: 2018/08/26 17:06
 */
type RawResponse struct {
	Result json.RawMessage   `json:"result"`
	Error  *Error 			 `json:"error"`
}

/*
Description:
GetResult checks whether the unmarshaled response contains a non-nil error,
returning an unmarshaled btcjson.RPCError (or an unmarshaling error) if so.
If the response is not an error, the raw bytes of the request are
returned for further unmashaling into specific result types.
 * Author: architect.bian
 * Date: 2018/08/26 17:54
 */
func (r RawResponse) GetResult() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}
