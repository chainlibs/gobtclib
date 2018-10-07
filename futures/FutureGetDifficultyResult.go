package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureGetDifficultyResult is a future promise to deliver the result of a
GetDifficultyAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
type FutureGetDifficultyResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
proof-of-work difficulty as a multiple of the minimum difficulty.
 * Author: architect.bian
 * Date: 2018/09/14 17:36
 */
func (r FutureGetDifficultyResult) Receive() (float64, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal the result as a float64.
	var difficulty float64
	err = json.Unmarshal(res, &difficulty)
	if err != nil {
		return 0, err
	}
	return difficulty, nil
}
