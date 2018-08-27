package client

import (
	"net/http"
	"encoding/json"
)


/*
Description:
sendPostDetails houses an HTTP POST request to send to an RPC server as well
as the original JSON-RPC command and a channel to reply on when the server
responds with the result.
 * Author: architect.bian
 * Date: 2018/08/26 15:42
 */
type sendPostDetails struct { //TODO? RequestDetail PostDetail?
	httpRequest *http.Request
	jsonRequest *jsonRequest
}

/*
Description:
jsonRequest holds information about a json request that is used to properly
detect, interpret, and deliver a reply to it.
 * Author: architect.bian
 * Date: 2018/08/26 15:43
 */
type jsonRequest struct { //TODO JsonDetail?JsonRequestDetail?
	id             uint64
	method         string
	cmd            interface{}
	marshalledJSON []byte
	responseChan   chan *response
}

/*
Description:
response is the raw bytes of a JSON-RPC result, or the error if the response
error object was non-null.
 * Author: architect.bian
 * Date: 2018/08/26 15:44
 */
type response struct {
	result []byte
	err    error
}


/*
Description:
rawResponse is a partially-unmarshaled JSON-RPC response.  For this
to be valid (according to JSON-RPC 1.0 spec), ID may not be nil.
 * Author: architect.bian
 * Date: 2018/08/26 17:06
 */
type rawResponse struct {
	Result json.RawMessage   `json:"result"`
	Error  *Error `json:"error"`
}

/*
Description:
result checks whether the unmarshaled response contains a non-nil error,
returning an unmarshaled btcjson.RPCError (or an unmarshaling error) if so.
If the response is not an error, the raw bytes of the request are
returned for further unmashaling into specific result types.
 * Author: architect.bian
 * Date: 2018/08/26 17:54
 */
func (r rawResponse) result() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}
