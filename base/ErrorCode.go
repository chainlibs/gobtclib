package base

/*
Description:
RPCErrorCode represents an error code to be used as a part of an RPCError
which is in turn used in a JSON-RPC JRPCResponse object.

A specific type is used to help ensure the wrong errors aren't used.
 * Author: architect.bian
 * Date: 2018/10/08 12:21
 */
type ErrorCode int

// These constants are used to identify a specific RuleError.
const (

	// ErrInvalidType indicates a type was passed that is not the required type.
	ErrInvalidTypeCode ErrorCode = iota

	//// ErrDuplicateMethodCode indicates a command with the specified method
	//// already exists.
	//ErrDuplicateMethodCode

)

