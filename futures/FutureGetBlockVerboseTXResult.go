package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetBlockVerboseTXResult is a future promise to deliver the result of a
GetBlockVerboseTXAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
type FutureGetBlockVerboseTXResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the data
structure from the server with information about the requested block.
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
func (r FutureGetBlockVerboseTXResult) Receive() (*results.GetBlockVerboseTXResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a BlockResult.
	var blockResult results.GetBlockVerboseTXResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil
}

