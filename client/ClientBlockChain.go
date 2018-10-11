package client

import (
	"github.com/chainlibs/gobtclib/results"
	"github.com/chainlibs/gobtclib/futures"
	"github.com/chainlibs/gobtclib/utils"
	"encoding/hex"
)

/*
Description:
GetBestBlockHashAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetBestBlockHash for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 13:18
 */
func (c *Client) GetBestBlockHashAsync() futures.FutureGetBestBlockHashResult {
	cmd := NewCommand("getbestblockhash")
	return c.sendCmd(cmd)
}

/*
Description:
GetBestBlockHash returns the hash of the best block in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:15
 */
func (c *Client) GetBestBlockHash() (*results.Hash, error) {
	return c.GetBestBlockHashAsync().Receive()
}

/*
Description:
GetBlockCountAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockCount for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 12:49
 */
func (c *Client) GetBlockCountAsync() futures.FutureGetBlockCountResult {
	cmd := NewCommand("getblockcount")
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockCount returns the number of blocks in the longest block chain.
 * Author: architect.bian
 * Date: 2018/09/14 12:46
 */
func (c *Client) GetBlockCount() (int64, error) {
	return c.GetBlockCountAsync().Receive()
}

/*
Description:
GetBlockHashAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHash for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/15 15:35
 */
func (c *Client) GetBlockHashAsync(blockHeight int64) futures.FutureGetBlockHashResult {
	cmd := NewCommand("getblockhash", blockHeight)
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHash returns the hash of the block in the best block chain at the given height.
 * Author: architect.bian
 * Date: 2018/09/15 15:34
 */
func (c *Client) GetBlockHash(blockHeight int64) (*results.Hash, error) {
	return c.GetBlockHashAsync(blockHeight).Receive()
}

/*
Description:
GetBlockHeaderAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockHeader for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 15:23
 */
func (c *Client) GetBlockHeaderAsync(blockHash *results.Hash, verbose *bool) futures.FutureResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblockheader", hash, verbose)
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockHeaderBytes returns the bytes of blockheader from the server given its hash.
See GetBlockHeader to retrieve a data structure with information about the block instead.
 * Author: architect.bian
 * Date: 2018/09/17 15:22
 */
func (c *Client) GetBlockHeaderBytes(blockHash *results.Hash) (*[]byte, error) {
	future := c.GetBlockHeaderAsync(blockHash, utils.Basis.Bool(false))
	result, err := futures.ReceiveFuture(future)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

/*
Description:
GetBlockHeader returns the blockheader from the server given its hash.
See GetBlockHeader to retrieve a data structure with information about the
block instead.
 * Author: architect.bian
 * Date: 2018/09/17 15:22
 */
func (c *Client) GetBlockHeader(blockHash *results.Hash) (*results.GetBlockHeaderResult, error) {
	return futures.FutureGetBlockHeaderResult(c.GetBlockHeaderAsync(blockHash, utils.Basis.Bool(true))).Receive()
}

/*
Description:
GetBlockChainInfoAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function
on the returned instance.

See GetBlockChainInfo for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 14:49
 */
func (c *Client) GetBlockChainInfoAsync() futures.FutureGetBlockChainInfoResult {
	cmd := NewCommand("getblockchaininfo")
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockChainInfo returns information related to the processing state of
various chain-specific details such as the current difficulty from the tip
of the main chain.
 * Author: architect.bian
 * Date: 2018/09/17 14:47
 */
func (c *Client) GetBlockChainInfo() (*results.GetBlockChainInfoResult, error) {
	return c.GetBlockChainInfoAsync().Receive()
}

/*
Description:
GetBlockAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlock for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 16:09
 */
func (c *Client) GetBlockAsync(blockHash *results.Hash, verbosity *int) futures.FutureResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}
	cmd := NewCommand("getblock", hash, verbosity)
	return c.sendCmd(cmd)
}

/*
Description:
GetBlock returns a raw block from the server given its hash.

See GetBlockVerbose to retrieve a data structure with information about the
block instead.
 * Author: architect.bian
 * Date: 2018/09/17 16:09
 */
func (c *Client) GetBlockBytes(blockHash *results.Hash) (*[]byte, error) {
	future := c.GetBlockAsync(blockHash, utils.Basis.Int(0))
	result, err := futures.ReceiveFuture(future)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

/*
Description:
GetBlock returns a raw block from the server given its hash.
See GetBlockVerbose to retrieve a data structure with information about the block instead.
* Author: architect.bian
* Date: 2018/09/17 16:09
*/
func (c *Client) GetBlock(blockHash *results.Hash) (*results.GetBlockResult, error) {
	future := futures.FutureGetBlockResult(c.GetBlockAsync(blockHash, utils.Basis.Int(1)))
	return future.Receive()
}

/*
Description:
GetBlockVerbose returns a data structure from the server with information
about a block given its hash.

See GetBlockVerboseTx to retrieve transaction data structures as well.
See GetBlock to retrieve a raw block instead.
 * Author: architect.bian
 * Date: 2018/09/17 16:21
 */
func (c *Client) GetBlockVerboseTX(blockHash *results.Hash) (*results.GetBlockVerboseTXResult, error) {
	future := futures.FutureGetBlockVerboseTXResult(c.GetBlockAsync(blockHash, utils.Basis.Int(2)))
	return future.Receive()
}

/*
Description:
GetBlockStatsAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetBlockStats for more details.
 * Author: architect.bian
 * Date: 2018/09/17 16:09
 */
func (c *Client) GetBlockStatsAsync(heightOrHash interface{}, stats *[]string) futures.FutureResult {
	cmd := NewCommand("getblockstats", heightOrHash, stats)
	return c.sendCmd(cmd)
}

/*
Description:
GetBlockStats Compute per block statistics for a given window. All amounts are in satoshis.
 * Author: architect.bian
 * Date: 2018/10/10 15:17
 */
func (c *Client) GetBlockStats(heightOrHash interface{}) (*results.GetBlockStatsResult, error) {
	var result results.GetBlockStatsResult
	_, err := c.GetBlockStatsAsync(heightOrHash, nil).Receive() //TODO result is nil? need *?
	if err != nil {
		return nil, err
	}
	return &result, err
}

/*
Description:
GetBlockStats Compute per block statistics for a given window. All amounts are in satoshis.
 * Author: architect.bian
 * Date: 2018/10/10 15:17
 */
func (c *Client) GetBlockStatsEntire(heightOrHash interface{}, stats *[]string) (*map[interface{}]interface{}, error) {
	var result map[interface{}]interface{}
	_, err := c.GetBlockStatsAsync(heightOrHash, nil).Receive()
	if err != nil {
		return nil, err
	}
	return &result, err
}

/*
Description:
GetChainTipsAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function
on the returned instance.

See GetChainTips for more details.
 * Author: architect.bian
 * Date: 2018/09/17 14:49
 */
func (c *Client) GetChainTipsAsync() futures.FutureResult {
	cmd := NewCommand("getchaintips")
	return c.sendCmd(cmd)
}

/*
Description:
GetChainTips return information about all known tips in the block tree, including the main chain
as well as orphaned branches.
of the main chain.
 * Author: architect.bian
 * Date: 2018/09/17 14:47
 */
func (c *Client) GetChainTips() (*[]results.GetChainTipsResult, error) {
	future := futures.FutureGetChainTipsResult(c.GetChainTipsAsync())
	return future.Receive()
}

/*
Description:
GetChainTXStatsAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetChainTXStats for more details.
 * Author: architect.bian
 * Date: 2018/09/17 16:09
 */
func (c *Client) GetChainTXStatsAsync(nblocks int32, blockhash *results.Hash) futures.FutureResult {
	cmd := NewCommand("getchaintxstats", nblocks, blockhash.String())
	return c.sendCmd(cmd)
}

/*
Description:
GetChainTXStats Compute statistics about the total number and rate of transactions in the chain.
 * Author: architect.bian
 * Date: 2018/10/11 10:23
 */
func (c *Client) GetChainTXStats() (*results.GetChainTXStatsResult, error) {
	cmd := NewCommand("getchaintxstats")
	var err error
	result, err := futures.FutureGetChainTXStatsResult(c.sendCmd(cmd)).Receive()
	if err != nil {
		return nil, err
	}
	return result, err
}

/*
Description:
GetChainTXStatsEntire Compute statistics about the total number and rate of transactions in the chain.
Can specify nblocks blockhash
 * Author: architect.bian
 * Date: 2018/10/11 10:24
 */
func (c *Client) GetChainTXStatsEntire(nblocks int32, blockhash *results.Hash) (*results.GetChainTXStatsResult, error) {
	result, err := futures.FutureGetChainTXStatsResult(c.GetChainTXStatsAsync(nblocks, blockhash)).Receive()
	if err != nil {
		return nil, err
	}
	return result, err
}

/*
Description:
GetDifficultyAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetDifficulty for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/14 17:35
 */
func (c *Client) GetDifficultyAsync() futures.FutureGetDifficultyResult {
	cmd := NewCommand("getdifficulty")
	return c.sendCmd(cmd)
}

/*
Description:
GetDifficulty returns the proof-of-work difficulty as a multiple of the
minimum difficulty. The result is bits/params.PowLimitBits
 * Author: architect.bian
 * Date: 2018/09/14 17:34
 */
func (c *Client) GetDifficulty() (float64, error) {
	return c.GetDifficultyAsync().Receive()
}

/*
Description:
GetMempoolAncestorsAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetMempoolAncestors for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:19
 */
func (c *Client) GetMempoolAncestorsAsync(txid results.Hash, verbose bool) futures.FutureResult {
	cmd := NewCommand("getmempoolancestors", txid.String(), verbose)
	return c.sendCmd(cmd)
}

/*
Description: 
GetMempoolAncestors If txid is in the mempool, returns all in-mempool ancestors.
 * Author: architect.bian
 * Date: 2018/10/11 11:41
 */
func (c *Client) GetMempoolAncestors(txid results.Hash) (*[]string, error) {
	var result []string
	_, err := c.GetMempoolAncestorsAsync(txid, false).Receive()
	if err != nil {
		return nil, err
	}
	return &result, err
}

/*
Description:
GetMempoolAncestorsVerbose If txid is in the mempool, returns all in-mempool ancestors.
verbose is true for a json object.
 * Author: architect.bian
 * Date: 2018/10/11 11:42
 */
func (c *Client) GetMempoolAncestorsVerbose(txid results.Hash) (*map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := c.GetMempoolAncestorsAsync(txid, true).Receive()
	if err != nil {
		return nil, err
	}
	return &result, err
}

/*
Description:
GetMempoolDescendantsAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetMempoolDescendants for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:19
 */
func (c *Client) GetMempoolDescendantsAsync(txid results.Hash, verbose bool) futures.FutureResult {
	cmd := NewCommand("getmempooldescendants", txid.String(), verbose)
	return c.sendCmd(cmd)
}

/*
Description:
GetMempoolDescendants If txid is in the mempool, returns all in-mempool descendants.
 * Author: architect.bian
 * Date: 2018/10/11 11:49
 */
func (c *Client) GetMempoolDescendants(txid results.Hash) (*[]string, error) {
	var result []string
	_, err := c.GetMempoolDescendantsAsync(txid, false).Receive()
	if err != nil {
		return nil, err
	}
	return &result, err
}

/*
Description:
GetMempoolDescendantsVerbose If txid is in the mempool, returns all in-mempool descendants.
 * Author: architect.bian
 * Date: 2018/10/11 11:49
 */
func (c *Client) GetMempoolDescendantsVerbose(txid results.Hash) (*map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := c.GetMempoolDescendantsAsync(txid, true).Receive()
	if err != nil {
		return nil, err
	}
	return &result, err
}

/*
Description:
GetMempoolEntryAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetMempoolEntry for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 20:23
 */
func (c *Client) GetMempoolEntryAsync(txHash string) futures.FutureGetMempoolEntryResult {
	cmd := NewCommand("getmempoolentry", txHash)
	return c.sendCmd(cmd)
}

/*
Description:
GetMempoolEntry returns a data structure with information about the
transaction in the memory pool given its hash.
 * Author: architect.bian
 * Date: 2018/09/17 20:23
 */
func (c *Client) GetMempoolEntry(txHash string) (*results.GetMempoolEntryResult, error) { //TODO txhash changed to hash?
	return c.GetMempoolEntryAsync(txHash).Receive()
}

/*
Description:
GetMempoolInfoAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetMempoolInfo for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:19
 */
func (c *Client) GetMempoolInfoAsync() futures.FutureResult {
	cmd := NewCommand("getmempoolinfo")
	return c.sendCmd(cmd)
}

/*
Description: 
GetMempoolInfo Returns details on the active state of the TX memory pool.
 * Author: architect.bian
 * Date: 2018/10/11 11:52
 */
func (c *Client) GetMempoolInfo() (*results.GetMempoolInfoResult, error) {
	result, err := futures.FutureGetMempoolInfoResult(c.GetMempoolInfoAsync()).Receive()
	if err != nil {
		return nil, err
	}
	return result, err
}

/*
Description:
GetRawMempoolAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetRawMempool for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 19:51
 */
func (c *Client) GetRawMempoolAsync() futures.FutureGetRawMempoolResult {
	cmd := NewCommand("getrawmempool", utils.Basis.Bool(false))
	return c.sendCmd(cmd)
}

/*
Description:
GetRawMempool returns the hashes of all transactions in the memory pool.

See GetRawMempoolVerbose to retrieve data structures with information about
the transactions instead.
 * Author: architect.bian
 * Date: 2018/09/17 19:50
 */
func (c *Client) GetRawMempool() ([]*results.Hash, error) {
	return c.GetRawMempoolAsync().Receive()
}

/*
Description:
GetRawMempoolVerboseAsync returns an instance of a type that can be used to
get the result of the RPC at some future time by invoking the Receive
function on the returned instance.

See GetRawMempoolVerbose for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 20:10
 */
func (c *Client) GetRawMempoolVerboseAsync() futures.FutureGetRawMempoolVerboseResult {
	cmd := NewCommand("getrawmempool", utils.Basis.Bool(true))
	return c.sendCmd(cmd)
}

/*
Description:
GetRawMempoolVerbose returns a map of transaction hashes to an associated
data structure with information about the transaction for all transactions in
the memory pool.

See GetRawMempool to retrieve only the transaction hashes instead.
 * Author: architect.bian
 * Date: 2018/09/17 20:09
 */
func (c *Client) GetRawMempoolVerbose() (map[string]results.GetRawMempoolVerboseResult, error) {
	return c.GetRawMempoolVerboseAsync().Receive()
}

/*
Description:
GetTxOutAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetTxOut for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 21:07
 */
func (c *Client) GetTxOutAsync(txHash *results.Hash, index uint32, mempool bool) futures.FutureGetTxOutResult {
	hash := ""
	if txHash != nil {
		hash = txHash.String()
	}
	cmd := NewCommand("gettxout", hash, index, &mempool)
	return c.sendCmd(cmd)
}

/*
Description:
GetTxOut returns the transaction output info if it's unspent and
nil, otherwise.
 * Author: architect.bian
 * Date: 2018/09/17 21:07
 */
func (c *Client) GetTxOut(txHash *results.Hash, index uint32, mempool bool) (*results.GetTxOutResult, error) {
	return c.GetTxOutAsync(txHash, index, mempool).Receive()
}

/*
Description:
GetTXOutProofAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetTXOutProof for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:20
 */
func (c *Client) GetTXOutProofAsync(txids []string, blockhash results.Hash) futures.FutureResult {
	cmd := NewCommand("gettxoutproof", txids, blockhash)
	return c.sendCmd(cmd)
}

func (c *Client) GetTXOutProof(txids ...string) (*[]byte, error) {
	cmd := NewCommand("gettxoutproof", txids)
	result, err := futures.ReceiveFuture(c.sendCmd(cmd))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetTXOutProofEntire(txids []string, blockhash results.Hash) (*[]byte, error) {
	future := c.GetTXOutProofAsync(txids, blockhash)
	result, err := futures.ReceiveFuture(future)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

/*
Description:
GetTXOutSetInfoAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See GetTXOutSetInfo for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:20
 */
func (c *Client) GetTXOutSetInfoAsync() futures.FutureResult {
	cmd := NewCommand("gettxoutsetinfo")
	return c.sendCmd(cmd)
}

func (c *Client) GetTXOutSetInfo() (*results.GetTXOutSetInfoResult, error) {
	result, err := futures.FutureGetTXOutSetInfoResult(c.GetTXOutSetInfoAsync()).Receive()
	if err != nil {
		return nil, err
	}
	return result, err
}

/*
Description:
PreciousBlockAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See PreciousBlock for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:20
 */
func (c *Client) PreciousBlockAsync(blockhash results.Hash) futures.FutureResult {
	cmd := NewCommand("preciousblock", blockhash.String())
	return c.sendCmd(cmd)
}

/*
Description: 
PreciousBlock Treats a block as if it were received before others with the same work.
 * Author: architect.bian
 * Date: 2018/10/11 12:33
 */
func (c *Client) PreciousBlock(blockhash results.Hash) (error) {
	future := c.PreciousBlockAsync(blockhash)
	_, err := futures.ReceiveFuture(future)
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
PruneBlockchainAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See PruneBlockchain for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:20
 */
func (c *Client) PruneBlockchainAsync(height int32) futures.FutureResult {
	cmd := NewCommand("pruneblockchain", height)
	return c.sendCmd(cmd)
}

/*
Description: 
PruneBlockchain prune blocks
 * Author: architect.bian
 * Date: 2018/10/11 12:33
 */
func (c *Client) PruneBlockchain(height int32) (int32, error) {
	var result int32
	_, err := c.PruneBlockchainAsync(height).Receive()
	if err != nil {
		return -1, err
	}
	return result, nil
}

/*
Description:
SaveMempoolAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See SaveMempool for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:20
 */
func (c *Client) SaveMempoolAsync() futures.FutureResult {
	cmd := NewCommand("savemempool")
	return c.sendCmd(cmd)
}

/*
Description:
SaveMempool Dumps the mempool to disk. It will fail until the previous dump is fully loaded.
 * Author: architect.bian
 * Date: 2018/10/11 12:32
 */
func (c *Client) SaveMempool() (error) {
	_, err := futures.ReceiveFuture(c.SaveMempoolAsync())
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
VerifyChainAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See VerifyChain for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 20:54
 */
func (c *Client) VerifyChainAsync() futures.FutureVerifyChainResult {
	cmd := NewCommand("verifychain")
	return c.sendCmd(cmd)
}

/*
Description:
VerifyChain requests the server to verify the block chain database using
the default check level and number of blocks to verify.

See VerifyChainLevel and VerifyChainBlocks to override the defaults.
 * Author: architect.bian
 * Date: 2018/09/17 20:54
 */
func (c *Client) VerifyChain() (bool, error) {
	return c.VerifyChainAsync().Receive()
}

/*
Description:
VerifyTXOutProofAsync returns an instance of a type that can be used to get the
result of the RPC at some future time by invoking the Receive function on the
returned instance.

See VerifyTXOutProof for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/10/11 10:20
 */
func (c *Client) VerifyTXOutProofAsync(proof []byte) futures.FutureResult {
	dst := make([]byte, hex.EncodedLen(len(proof)))
	hex.Encode(dst, proof)
	cmd := NewCommand("verifytxoutproof", string(dst))
	return c.sendCmd(cmd)
}

func (c *Client) VerifyTXOutProof(proof []byte) (*[]string, error) {
	var result []string
	_, err := c.VerifyTXOutProofAsync(proof).Receive()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

/*
Description:
VerifyChainLevelAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See VerifyChainLevel for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 21:00
 */
func (c *Client) VerifyChainLevelAsync(checkLevel int32) futures.FutureVerifyChainResult {
	cmd := NewCommand("verifychain", &checkLevel)
	return c.sendCmd(cmd)
}

/*
Description:
VerifyChainLevel requests the server to verify the block chain database using
the passed check level and default number of blocks to verify.

The check level controls how thorough the verification is with higher numbers
increasing the amount of checks done as consequently how long the
verification takes.

See VerifyChain to use the default check level and VerifyChainBlocks to
override the number of blocks to verify.
 * Author: architect.bian
 * Date: 2018/09/17 21:00
 */
func (c *Client) VerifyChainLevel(checkLevel int32) (bool, error) {
	return c.VerifyChainLevelAsync(checkLevel).Receive()
}

/*
Description:
VerifyChainBlocksAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See VerifyChainBlocks for the blocking version and more details.
 * Author: architect.bian
 * Date: 2018/09/17 21:00
 */
func (c *Client) VerifyChainBlocksAsync(checkLevel, numBlocks int32) futures.FutureVerifyChainResult {
	cmd := NewCommand("verifychain", &checkLevel, &numBlocks)
	return c.sendCmd(cmd)
}

/*
Description:
VerifyChainBlocks requests the server to verify the block chain database
using the passed check level and number of blocks to verify.

The check level controls how thorough the verification is with higher numbers
increasing the amount of checks done as consequently how long the
verification takes.

The number of blocks refers to the number of blocks from the end of the
current longest chain.

See VerifyChain and VerifyChainLevel to use defaults.
 * Author: architect.bian
 * Date: 2018/09/17 20:59
 */
func (c *Client) VerifyChainBlocks(checkLevel, numBlocks int32) (bool, error) {
	return c.VerifyChainBlocksAsync(checkLevel, numBlocks).Receive()
}
