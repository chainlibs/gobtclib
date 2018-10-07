package futures
import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetBlockHeaderVerboseResult is a future promise to deliver the result of a
GetBlockAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/15 17:52
 */
type FutureGetBlockHeaderVerboseResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
data structure of the blockheader requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/09/15 17:52
 */
func (r FutureGetBlockHeaderVerboseResult) Receive() (*results.GetBlockHeaderVerboseResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var bh results.GetBlockHeaderVerboseResult
	err = json.Unmarshal(res, &bh)
	if err != nil {
		return nil, err
	}

	return &bh, nil
}

