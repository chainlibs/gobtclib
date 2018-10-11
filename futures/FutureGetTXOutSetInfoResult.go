package futures
import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetBlockChainInfoResult is a promise to deliver the result of a
GetBlockChainInfoAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 14:51
 */
type FutureGetTXOutSetInfoResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns chain info
result provided by the server.
 * Author: architect.bian
 * Date: 2018/09/17 14:52
 */
func (r FutureGetTXOutSetInfoResult) Receive() (*results.GetTXOutSetInfoResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result results.GetTXOutSetInfoResult
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

