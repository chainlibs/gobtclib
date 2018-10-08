package base

import (
	"encoding/json"
	"fmt"
)

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
	if !isValidIDType(id) {
		return nil, NewError(ErrInvalidTypeCode, fmt.Sprintf("the id of type '%T' is invalid", id))
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
isValidIDType checks that the ID field (which can go in any of the JSON-RPC
requests, responses, or notifications) is valid.  JSON-RPC 1.0 allows any
valid JSON type.  JSON-RPC 2.0 (which bitcoind follows for some parts) only
allows string, number, or null, so this function restricts the allowed types
to that list.  This function is only provided in case the caller is manually
marshalling for some reason.    The functions which accept an ID in this
package already call this function to ensure the provided id is valid.
 * Author: architect.bian
 * Date: 2018/08/26 17:23
 */
func isValidIDType(id interface{}) bool {
	switch id.(type) {
	case int, int8, int16, int32, int64,
	uint, uint8, uint16, uint32, uint64,
	float32, float64, string, nil:
		return true
	default:
		return false
	}
}
