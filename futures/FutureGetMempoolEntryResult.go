package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetMempoolEntryResult is a future promise to deliver the result of a
GetMempoolEntryAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 20:24
 */
type FutureGetMempoolEntryResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns a data
structure with information about the transaction in the memory pool given
its hash.
 * Author: architect.bian
 * Date: 2018/09/17 20:24
 */
func (r FutureGetMempoolEntryResult) Receive() (*results.GetMempoolEntryResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an array of strings.
	var mempoolEntryResult results.GetMempoolEntryResult
	err = json.Unmarshal(res, &mempoolEntryResult)
	if err != nil {
		return nil, err
	}

	return &mempoolEntryResult, nil
}

