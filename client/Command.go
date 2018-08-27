package client

import (
	"reflect"
	"fmt"
	"sync"
	"encoding/json"
	"strings"
	"strconv"
)

// UsageFlag define flags that specify additional properties about the
// circumstances under which a command can be used.
type UsageFlag uint32

const (
	// UFWalletOnly indicates that the command can only be used with an RPC
	// server that supports wallet commands.
	UFWalletOnly UsageFlag = 1 << iota

	// UFWebsocketOnly indicates that the command can only be used when
	// communicating with an RPC server over websockets.  This typically
	// applies to notifications and notification registration functions
	// since neiher makes since when using a single-shot HTTP-POST request.
	UFWebsocketOnly

	// UFNotification indicates that the command is actually a notification.
	// This means when it is marshalled, the ID must be nil.
	UFNotification

	// highestUsageFlagBit is the maximum usage flag bit and is used in the
	// stringer and tests to ensure all of the above constants have been
	// tested.
	highestUsageFlagBit
)

// Map of UsageFlag values back to their constant names for pretty printing.
var usageFlagStrings = map[UsageFlag]string{
	UFWalletOnly:    "UFWalletOnly",
	UFWebsocketOnly: "UFWebsocketOnly",
	UFNotification:  "UFNotification",
}

// String returns the UsageFlag in human-readable form.
func (fl UsageFlag) String() string {
	// No flags are set.
	if fl == 0 {
		return "0x0"
	}

	// Add individual bit flags.
	s := ""
	for flag := UFWalletOnly; flag < highestUsageFlagBit; flag <<= 1 {
		if fl&flag == flag {
			s += usageFlagStrings[flag] + "|"
			fl -= flag
		}
	}

	// Add remaining value as raw hex.
	s = strings.TrimRight(s, "|")
	if fl != 0 {
		s += "|0x" + strconv.FormatUint(uint64(fl), 16)
	}
	s = strings.TrimLeft(s, "|")
	return s
}

// methodInfo keeps track of information about each registered method such as
// the parameter information.
type methodInfo struct {
	maxParams    int
	numReqParams int
	numOptParams int
	defaults     map[int]reflect.Value
	flags        UsageFlag
	usage        string
}

var (
	// These fields are used to map the registered types to method names.
	registerLock         sync.RWMutex
	methodToConcreteType = make(map[string]reflect.Type)
	methodToInfo         = make(map[string]methodInfo)
	concreteTypeToMethod = make(map[reflect.Type]string)
)

/*
Description:
getCmdMethod returns the method for the passed command.  The provided command
type must be a registered type.  All commands provided by this package are registered by default.
 * Author: architect.bian
 * Date: 2018/08/26 18:48
 */
func getCmdMethod(cmd interface{}) (string, error) {
	// Look up the cmd type and error out if not registered.
	rt := reflect.TypeOf(cmd)
	registerLock.RLock()
	method, ok := concreteTypeToMethod[rt]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registered", method)
		return "", makeError(ErrUnregisteredMethod, str)
	}

	return method, nil
}

/*
Description:
MarshalCmd marshals the passed command to a JSON-RPC request byte slice that
is suitable for transmission to an RPC server.  The provided command type
must be a registered type.  All commands provided by this package are
registered by default.
 * Author: architect.bian
 * Date: 2018/08/26 19:14
 */
func MarshalCmd(id interface{}, cmd interface{}) ([]byte, error) {
	// Look up the cmd type and error out if not registered.
	rt := reflect.TypeOf(cmd)
	registerLock.RLock()
	method, ok := concreteTypeToMethod[rt]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registered", method)
		return nil, makeError(ErrUnregisteredMethod, str)
	}

	// The provided command must not be nil.
	rv := reflect.ValueOf(cmd)
	if rv.IsNil() {
		str := "the specified command is nil"
		return nil, makeError(ErrInvalidType, str)
	}

	// Create a slice of interface values in the order of the struct fields
	// while respecting pointer fields as optional params and only adding
	// them if they are non-nil.
	params := makeParams(rt.Elem(), rv.Elem())

	// Generate and marshal the final JSON-RPC request.
	jsonRPC, err := NewRequest(id, method, params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonRPC)
}

/*
Description:
makeParams creates a slice of interface values for the given struct.
 * Author: architect.bian
 * Date: 2018/08/26 19:02
 */
func makeParams(rt reflect.Type, rv reflect.Value) []interface{} {
	numFields := rt.NumField()
	params := make([]interface{}, 0, numFields)
	for i := 0; i < numFields; i++ {
		rtf := rt.Field(i)
		rvf := rv.Field(i)
		if rtf.Type.Kind() == reflect.Ptr {
			if rvf.IsNil() {
				break
			}
			params = append(params, rvf.Elem())
		}
		params = append(params, rvf.Interface())
	}

	return params
}