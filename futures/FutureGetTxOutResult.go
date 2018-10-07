package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetTxOutResult is a future promise to deliver the result of a
GetTxOutAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 21:08
 */
type FutureGetTxOutResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns a
transaction given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 21:08
 */
func (r FutureGetTxOutResult) Receive() (*results.GetTxOutResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// take care of the special case where the output has been spent already
	// it should return the string "null"
	if string(res) == "null" {
		return nil, nil
	}

	// Unmarshal result as an gettxout result object.
	var txOutInfo *results.GetTxOutResult
	err = json.Unmarshal(res, &txOutInfo)
	if err != nil {
		return nil, err
	}

	return txOutInfo, nil
}

