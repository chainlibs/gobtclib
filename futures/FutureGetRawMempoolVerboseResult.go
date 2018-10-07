package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetRawMempoolVerboseResult is a future promise to deliver the result of
a GetRawMempoolVerboseAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 20:11
 */
type FutureGetRawMempoolVerboseResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns a map of
transaction hashes to an associated data structure with information about the
transaction for all transactions in the memory pool.
 * Author: architect.bian
 * Date: 2018/09/17 20:11
 */
func (r FutureGetRawMempoolVerboseResult) Receive() (map[string]results.GetRawMempoolVerboseResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a map of strings (tx shas) to their detailed results.
	var mempoolItems map[string]results.GetRawMempoolVerboseResult
	err = json.Unmarshal(res, &mempoolItems)
	if err != nil {
		return nil, err
	}
	return mempoolItems, nil
}

