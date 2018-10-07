package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
)

/*
Description:
FutureVerifyChainResult is a future promise to deliver the result of a
VerifyChainAsync, VerifyChainLevelAsyncRPC, or VerifyChainBlocksAsync
invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 20:54
 */
type FutureVerifyChainResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns whether
or not the chain verified based on the check level and number of blocks
to verify specified in the original call.
 * Author: architect.bian
 * Date: 2018/09/17 20:54
 */
func (r FutureVerifyChainResult) Receive() (bool, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return false, err
	}

	// Unmarshal the result as a boolean.
	var verified bool
	err = json.Unmarshal(res, &verified)
	if err != nil {
		return false, err
	}
	return verified, nil
}

