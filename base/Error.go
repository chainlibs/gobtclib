package base

import (
	"errors"
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
RPCErrorCode represents an error code to be used as a part of an RPCError
which is in turn used in a JSON-RPC JRPCResponse object.

A specific type is used to help ensure the wrong errors aren't used.
 * Author: architect.bian
 * Date: 2018/10/08 12:21
 */
type ErrorCode int

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












var (
	// ErrInvalidAuth is an error to describe the condition where the client
	// is either unable to authenticate or the specified endpoint is
	// incorrect.
	ErrInvalidAuth = errors.New("authentication failure")

	// ErrInvalidEndpoint is an error to describe the condition where the
	// websocket handshake failed with the specified endpoint.
	ErrInvalidEndpoint = errors.New("the endpoint either does not support " +
		"websockets or does not exist")

	// ErrClientNotConnected is an error to describe the condition where a
	// websocket client has been created, but the connection was never
	// established.  This condition differs from ErrClientDisconnect, which
	// represents an established connection that was lost.
	ErrClientNotConnected = errors.New("the client was never connected")

	// ErrClientDisconnect is an error to describe the condition where the
	// client has been disconnected from the RPC server.  When the
	// DisableAutoReconnect option is not set, any outstanding futures
	// when a client disconnect occurs will return this error as will
	// any new requests.
	ErrClientDisconnect = errors.New("the client has been disconnected")

	// ErrClientShutdown is an error to describe the condition where the
	// client is either already shutdown, or in the process of shutting
	// down.  Any outstanding futures when a client shutdown occurs will
	// return this error as will any new requests.
	ErrClientShutdown = errors.New("the client has been shutdown")

	// ErrNotWebsocketClient is an error to describe the condition of
	// calling a Client method intended for a websocket client when the
	// client has been configured to run in HTTP POST mode instead.
	ErrNotWebsocketClient = errors.New("client is not configured for " +
		"websockets")

	// ErrClientAlreadyConnected is an error to describe the condition where
	// a new client connection cannot be established due to a websocket
	// client having already connected to the RPC server.
	ErrClientAlreadyConnected = errors.New("websocket client has already " +
		"connected")
)

// These constants are used to identify a specific RuleError.
const (
	// ErrDuplicateMethod indicates a command with the specified method
	// already exists.
	ErrDuplicateMethod ErrorCode = iota

	// ErrInvalidUsageFlags indicates one or more unrecognized flag bits
	// were specified.
	ErrInvalidUsageFlags

	// ErrInvalidType indicates a type was passed that is not the required
	// type.
	ErrInvalidType

	// ErrEmbeddedType indicates the provided command struct contains an
	// embedded type which is not not supported.
	ErrEmbeddedType

	// ErrUnexportedField indiciates the provided command struct contains an
	// unexported field which is not supported.
	ErrUnexportedField

	// ErrUnsupportedFieldType indicates the type of a field in the provided
	// command struct is not one of the supported types.
	ErrUnsupportedFieldType

	// ErrNonOptionalField indicates a non-optional field was specified
	// after an optional field.
	ErrNonOptionalField

	// ErrNonOptionalDefault indicates a 'jsonrpcdefault' struct tag was
	// specified for a non-optional field.
	ErrNonOptionalDefault

	// ErrMismatchedDefault indicates a 'jsonrpcdefault' struct tag contains
	// a value that doesn't match the type of the field.
	ErrMismatchedDefault

	// ErrUnregisteredMethod indicates a method was specified that has not
	// been registered.
	ErrUnregisteredMethod

	// ErrMissingDescription indicates a description required to generate
	// help is missing.
	ErrMissingDescription

	// ErrNumParams inidcates the number of params supplied do not
	// match the requirements of the associated command.
	ErrNumParams

	// numErrorCodes is the maximum error code number used in tests.
	numErrorCodes
)
