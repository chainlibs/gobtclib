package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
	"encoding/hex"
	"bytes"
)

/*
Description:
FutureGetBlockResult is a future promise to deliver the result of a
GetBlockAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 16:11
 */
type FutureGetBlockResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the raw
block requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 16:12
 */
func (r FutureGetBlockResult) Receive() (*results.MsgBlock, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var blockHex string
	err = json.Unmarshal(res, &blockHex)
	if err != nil {
		return nil, err
	}

	// Decode the serialized block hex to raw bytes.
	serializedBlock, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the block and return it.
	var msgBlock results.MsgBlock
	err = msgBlock.Deserialize(bytes.NewReader(serializedBlock))
	if err != nil {
		return nil, err
	}
	return &msgBlock, nil
}

