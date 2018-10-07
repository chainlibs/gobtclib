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
FutureGetBlockHeaderResult is a future promise to deliver the result of a
GetBlockHeaderAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 15:23
 */
type FutureGetBlockHeaderResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
blockheader requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 15:23
 */
func (r FutureGetBlockHeaderResult) Receive() (*results.BlockHeader, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var bhHex string
	err = json.Unmarshal(res, &bhHex)
	if err != nil {
		return nil, err
	}

	serializedBH, err := hex.DecodeString(bhHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the blockheader and return it.
	var bh results.BlockHeader
	err = bh.Deserialize(bytes.NewReader(serializedBH))
	if err != nil {
		return nil, err
	}

	return &bh, err
}

