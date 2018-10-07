package futures

import (
	"encoding/json"
	"github.com/chainlibs/gobtclib/base"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
FutureGetRawMempoolResult is a future promise to deliver the result of a
GetRawMempoolAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 19:54
 */
type FutureGetRawMempoolResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the hashes
of all transactions in the memory pool.
 * Author: architect.bian
 * Date: 2018/09/17 19:54
 */
func (r FutureGetRawMempoolResult) Receive() ([]*results.Hash, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an array of strings.
	var txHashStrs []string
	err = json.Unmarshal(res, &txHashStrs)
	if err != nil {
		return nil, err
	}

	// Create a slice of ShaHash arrays from the string slice.
	txHashes := make([]*results.Hash, 0, len(txHashStrs))
	for _, hashStr := range txHashStrs {
		txHash, err := results.NewHashFromStr(hashStr)
		if err != nil {
			return nil, err
		}
		txHashes = append(txHashes, txHash)
	}

	return txHashes, nil
}

