package base

import (
	"fmt"
)

/*
Description:
NewError creates an Error given a set of arguments.
Constructs and returns a new JSON-RPC error that is suitable
for use in a JSON-RPC JRPCResponse object.
 * Author: architect.bian
 * Date: 2018/10/08 12:09
 */
func NewError(c ErrorCode, msg string) Error {
	return Error{Code: c, Message: msg}
}

/*
Description:
NewErrorF create an Error given an ErrorCode and a format specifier then returns the error.
 * Author: architect.bian
 * Date: 2018/10/08 12:17
 */
func NewErrorF(c ErrorCode, format string, a ...interface{}) Error {
	return NewError(c, fmt.Sprintf(format, a...))
}

/*
Description:
RPCError represents an error that is used as a part of a JSON-RPC JRPCResponse object.
 * Author: architect.bian
 * Date: 2018/10/08 12:21
 */
type Error struct {
	Code    ErrorCode	 `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}

/*
Description:
Guarantee RPCError satisifies the builtin error interface.
 * Author: architect.bian
 * Date: 2018/10/08 12:22
 */
var _, _ error = Error{}, (*Error)(nil)

/*
Description:
Error returns a string describing the RPC error.
This satisifies the builtin error interface.
 * Author: architect.bian
 * Date: 2018/10/08 12:23
 */
func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
