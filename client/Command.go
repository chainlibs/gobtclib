package client

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
NewCommand create a new command with cmd and args.
arguments:
name, the command name.
args, the arguments for current command
 * Author: architect.bian
 * Date: 2018/10/06 18:26
 */
func NewCommand(name string, args ...interface{}) *Command {
	return &Command{
		name: name,
		args: args,
	}
}

/*
Description:
MarshalCmdToJRPC marshals the passed command to a JSON-RPC request byte slice that
is suitable for transmission to an RPC server.  The provided command type
must be a registered type.  All commands provided by this package are
registered by default.
 * Author: architect.bian
 * Date: 2018/08/26 19:14
 */
func MarshalCmdToJRPC(id uint64, cmd *Command) ([]byte, error) {
	// Generate and marshal the final JSON-RPC request.
	jsonRPC, err := base.NewJRPC(id, cmd.name, cmd.args)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonRPC)
}

/*
Description: 
Command represents a rpc command
 * Author: architect.bian
 * Date: 2018/10/06 18:28
 */
type Command struct {
	name string
	args []interface{}
}

func (c *Command) AddArgs(args ...interface{}) {
	c.args = append(c.args[:], args...)
}

func (c *Command) AddJsonArgs(args ...interface{}) error {
	slice := make([]interface{}, 0)
	for _, item := range args  {
		bytes, err := json.Marshal(item)
		if err != nil {
			return nil
		}
		slice = append(slice, string(bytes))
	}

	c.args = append(c.args[:], slice...)
	return nil
}
