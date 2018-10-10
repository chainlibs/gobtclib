package futures

import (
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
	"encoding/json"
)

/*
Description:
FutureGetChainTipsResult is a future promise to deliver the result of a
GetChainTipsAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/10/10 13:56
 */
type FutureGetChainTipsResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the data
structure from the server with information about the requested block.
 * Author: architect.bian
 * Date: 2018/10/10 13:57
 */
func (f FutureGetChainTipsResult) Receive() (*[]results.GetChainTipsResult, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return nil, err
	}

	var result []results.GetChainTipsResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}