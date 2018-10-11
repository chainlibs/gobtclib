package futures

import (
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
	"encoding/json"
	"fmt"
)

/*
Description:
FutureGetChainTXStatsResult is a future promise to deliver the result of a
GetChainTXStats RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/10/10 13:56
 */
type FutureGetChainTXStatsResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the data
structure from the server with information about the requested block.
 * Author: architect.bian
 * Date: 2018/10/10 13:57
 */
func (f FutureGetChainTXStatsResult) Receive() (*results.GetChainTXStatsResult, error) {
	res, err := ReceiveFuture(f)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))

	var result results.GetChainTXStatsResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}