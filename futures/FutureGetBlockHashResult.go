package futures
import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetBlockHashResult is a future promise to deliver the result of a
GetBlockHashAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/15 15:44
 */
type FutureGetBlockHashResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the hash of
the block in the best block chain at the given height.
 * Author: architect.bian
 * Date: 2018/09/15 15:45
 */
func (r FutureGetBlockHashResult) Receive() (*results.Hash, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a string-encoded sha.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return nil, err
	}
	return results.NewHashFromStr(txHashStr)
}
