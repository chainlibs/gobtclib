package base

import (
	"encoding/json"
)

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
RespRaw is a partially-unmarshaled JSON-RPC response.  For this
to be valid (according to JSON-RPC 1.0 spec), ID may not be nil.
 * Author: architect.bian
 * Date: 2018/08/26 17:06
 */
type RespRaw struct {
	Result json.RawMessage   `json:"result"`
	Error  *Error 			 `json:"error"`
}

/*
Description:
GetRaw checks whether the unmarshaled response contains a non-nil error,
returning an unmarshaled btcjson.RPCError (or an unmarshaling error) if so.
If the response is not an error, the raw bytes of the request are
returned for further unmashaling into specific result types.
 * Author: architect.bian
 * Date: 2018/08/26 17:54
 */
func (r RespRaw) GetRaw() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}
