package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetBlockResult is a future promise to deliver the result of a
GetBlockVerboseAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
type FutureGetBlockResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the data
structure from the server with information about the requested block.
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
func (f FutureGetBlockResult) Receive() (*results.GetBlockResult, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a BlockResult.
	var blockResult results.GetBlockResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil
}

