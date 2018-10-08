package base

import (
	"encoding/json"
	"fmt"
)

/*
Description:
IsValidIDType checks that the ID field (which can go in any of the JSON-RPC
requests, responses, or notifications) is valid.  JSON-RPC 1.0 allows any
valid JSON type.  JSON-RPC 2.0 (which bitcoind follows for some parts) only
allows string, number, or null, so this function restricts the allowed types
to that list.  This function is only provided in case the caller is manually
marshalling for some reason.    The functions which accept an ID in this
package already call this function to ensure the provided id is valid.
 * Author: architect.bian
 * Date: 2018/08/26 17:23
 */
func IsValidIDType(id interface{}) bool {
	switch id.(type) {
	case int, int8, int16, int32, int64,
	uint, uint8, uint16, uint32, uint64,
	float32, float64, string, nil:
		return true
	default:
		return false
	}
}

/*
Description:
JRPC is a type for raw JSON-RPC 1.0 requests.  The Method field identifies
the specific command type which in turns leads to different parameters.
Callers typically will not use this directly since this package provides a
statically typed command infrastructure which handles creation of these
requests, however this struct it being exported in case the caller wants to
construct raw requests for some reason.
 * Author: architect.bian
 * Date: 2018/08/26 17:27
 */
type JRPC struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ID      interface{}       `json:"id"`
}

/*
Description:
NewJRPC returns a new JSON-RPC 1.0 request object given the provided id,
method, and parameters.  The parameters are marshalled into a json.RawMessage
for the Params field of the returned request object.  This function is only
provided in case the caller wants to construct raw requests for some reason.

Typically callers will instead want to create a registered concrete command
type with the NewCmd or New<Foo>Cmd functions and call the MarshalCmd
function with that command to generate the marshalled JSON-RPC request.
 * Author: architect.bian
 * Date: 2018/08/26 17:29
 */
func NewJRPC(id interface{}, method string, params []interface{}) (*JRPC, error) {
	if !IsValidIDType(id) {
		return nil, NewError(ErrInvalidType, fmt.Sprintf("the id of type '%T' is invalid", id))
	}

	rawParams := make([]json.RawMessage, 0, len(params))
	for _, param := range params {
		marshalledParam, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		rawMessage := json.RawMessage(marshalledParam)
		rawParams = append(rawParams, rawMessage)
	}

	return &JRPC{
		Jsonrpc: "1.0",
		ID:      id,
		Method:  method,
		Params:  rawParams,
	}, nil
}

///*
//Description:
//JRPCResponse is the general form of a JSON-RPC response.  The type of the Result
//field varies from one command to the next, so it is implemented as an
//interface.  The ID field has to be a pointer for Go to put a null in it when empty.
// * Author: architect.bian
// * Date: 2018/08/26 17:35
// */
//type JRPCResponse struct {
//	Result json.RawMessage `json:"result"`
//	Error  *RPCError       `json:"error"`
//	ID     *interface{}    `json:"id"`
//}

///*
//Description:
//NewJRPCResponse returns a new JSON-RPC response object given the provided id,marshalled result, and RPC error.
//This function is only provided in case the caller wants to construct raw responses for some reason.
//
//Typically callers will instead want to create the fully marshalled JSON-RPC
//response to send over the wire with the MarshalJRPCResponse function.
// * Author: architect.bian
// * Date: 2018/08/26 17:36
// */
//func NewJRPCResponse(id interface{}, marshalledResult []byte, rpcErr *RPCError) (*JRPCResponse, error) {
//	if !IsValidIDType(id) {
//		return nil, makeInvalidIDError(ErrInvalidType, id)
//	}
//
//	pid := &id
//	return &JRPCResponse{
//		Result: marshalledResult,
//		Error:  rpcErr,
//		ID:     pid,
//	}, nil
//}
//
///*
//Description:
//MarshalJRPCResponse marshals the passed id, result, and RPCError to a JSON-RPC
//response byte slice that is suitable for transmission to a JSON-RPC client.
// * Author: architect.bian
// * Date: 2018/08/26 17:44
// */
//func MarshalJRPCResponse(id interface{}, result interface{}, rpcErr *RPCError) ([]byte, error) {
//	marshalledResult, err := json.Marshal(result)
//	if err != nil {
//		return nil, err
//	}
//	response, err := NewJRPCResponse(id, marshalledResult, rpcErr)
//	if err != nil {
//		return nil, err
//	}
//	return json.Marshal(&response)
//}

///*
//Description:
//NewRPCError constructs and returns a new JSON-RPC error that is suitable
//for use in a JSON-RPC JRPCResponse object.
// * Author: architect.bian
// * Date: 2018/08/26 17:46
// */
//func NewRPCError(code RPCErrorCode, message string) *RPCError {
//	return &RPCError{
//		Code:    code,
//		Message: message,
//	}
//}

///*
//Description:
//RPCErrorCode represents an error code to be used as a part of an RPCError
//which is in turn used in a JSON-RPC JRPCResponse object.
//
//A specific type is used to help ensure the wrong errors aren't used.
// * Author: architect.bian
// * Date: 2018/08/26 17:46
// */
//type RPCErrorCode int
//
///*
//Description:
//RPCError represents an error that is used as a part of a JSON-RPC JRPCResponse object.
// * Author: architect.bian
// * Date: 2018/08/26 17:16
// */
//type RPCError struct {
//	Code    RPCErrorCode `json:"code,omitempty"`
//	Message string       `json:"message,omitempty"`
//}
//
///*
//Description:
//Error returns a string describing the RPC error.  This satisifies the
//builtin error interface.
// * Author: architect.bian
// * Date: 2018/08/26 17:19
// */
//func (e RPCError) Error() string {
//	return fmt.Sprintf("%d: %s", e.Code, e.Message)
//}
//
//// Guarantee RPCError satisifies the builtin error interface.
//var _, _ error = RPCError{}, (*RPCError)(nil)
