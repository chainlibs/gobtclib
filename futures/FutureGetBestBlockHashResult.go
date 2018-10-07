package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetBestBlockHashResult is a future promise to deliver the result of a
GetBestBlockAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/14 13:20
 */
type FutureGetBestBlockHashResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the hash of
the best block in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:20
 */
func (r FutureGetBestBlockHashResult) Receive() (*results.Hash, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return nil, err
	}
	return results.NewHashFromStr(txHashStr)
}

