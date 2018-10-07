package futures

import (
	"encoding/json"
	"encoding/hex"
	"bytes"
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

/*
Description:
FutureGetBlockHeaderResult is a future promise to deliver the result of a
GetBlockHeaderAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 15:23
 */
type FutureGetBlockHeaderResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the
blockheader requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 15:23
 */
func (r FutureGetBlockHeaderResult) Receive() (*results.BlockHeader, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var bhHex string
	err = json.Unmarshal(res, &bhHex)
	if err != nil {
		return nil, err
	}

	serializedBH, err := hex.DecodeString(bhHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the blockheader and return it.
	var bh results.BlockHeader
	err = bh.Deserialize(bytes.NewReader(serializedBH))
	if err != nil {
		return nil, err
	}

	return &bh, err
}

/*
Description:
FutureGetBlockChainInfoResult is a promise to deliver the result of a
GetBlockChainInfoAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 14:51
 */
type FutureGetBlockChainInfoResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns chain info
result provided by the server.
 * Author: architect.bian
 * Date: 2018/09/17 14:52
 */
func (r FutureGetBlockChainInfoResult) Receive() (*results.GetBlockChainInfoResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	var chainInfo results.GetBlockChainInfoResult
	if err := json.Unmarshal(res, &chainInfo); err != nil {
		return nil, err
	}
	return &chainInfo, nil
}

/*
Description:
FutureGetBlockResult is a future promise to deliver the result of a
GetBlockAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 16:11
 */
type FutureGetBlockResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the raw
block requested from the server given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 16:12
 */
func (r FutureGetBlockResult) Receive() (*results.MsgBlock, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var blockHex string
	err = json.Unmarshal(res, &blockHex)
	if err != nil {
		return nil, err
	}

	// Decode the serialized block hex to raw bytes.
	serializedBlock, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, err
	}

	// Deserialize the block and return it.
	var msgBlock results.MsgBlock
	err = msgBlock.Deserialize(bytes.NewReader(serializedBlock))
	if err != nil {
		return nil, err
	}
	return &msgBlock, nil
}

/*
Description:
FutureGetBlockVerboseResult is a future promise to deliver the result of a
GetBlockVerboseAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
type FutureGetBlockVerboseResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the data
structure from the server with information about the requested block.
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
func (r FutureGetBlockVerboseResult) Receive() (*results.GetBlockVerboseResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a BlockResult.
	var blockResult results.GetBlockVerboseResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil
}

/*
Description:
FutureGetBlockVerboseTXResult is a future promise to deliver the result of a
GetBlockVerboseTXAsync RPC invocation (or an applicable error).
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
type FutureGetBlockVerboseTXResult chan *base.Response

/*
Description:
Receive waits for the response promised by the future and returns the data
structure from the server with information about the requested block.
 * Author: architect.bian
 * Date: 2018/09/17 16:26
 */
func (r FutureGetBlockVerboseTXResult) Receive() (*results.GetBlockVerboseTXResult, error) {
	res, err := ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the raw result into a BlockResult.
	var blockResult results.GetBlockVerboseTXResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}
	return &blockResult, nil
}

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
