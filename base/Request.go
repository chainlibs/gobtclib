package base

import (
	"net/http"
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
	Detail      *JsonDetail
}

/*
Description:
JsonDetail holds information about a json request that is used to properly
detect, interpret, and deliver a reply to it.
 * Author: architect.bian
 * Date: 2018/08/26 15:43
 */
type JsonDetail struct {
	Id             uint64
	Method         string
	Cmd            interface{}
	MarshalledJSON []byte
	ResponseChan   chan *Response // Generate the request and send it along with a channel to respond on.
}
