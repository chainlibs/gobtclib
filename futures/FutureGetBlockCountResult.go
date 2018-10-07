package futures
import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureGetBlockCountResult is a future promise to deliver the result of a
GetBlockCountAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/14 12:55
 */
type FutureGetBlockCountResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the number
of blocks in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 12:59
 */
func (r FutureGetBlockCountResult) Receive() (int64, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return -1, err
	}

	// Unmarshal the result as an int64.
	var count int64
	err = json.Unmarshal(res, &count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
