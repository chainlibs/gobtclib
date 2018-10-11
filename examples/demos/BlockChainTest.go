package demos

import (
	"github.com/gobasis/log"
	"github.com/chainlibs/gobtclib/results"
)

/*
Description:
a demo test of GetBlockCount, Get the current block count.
 * Author: architect.bian
 * Date: 2018/09/02 18:24
 */
func GetBlockCountTest() {
	blockCount, err := cli.GetBlockCount()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockCount", "count", blockCount)
}

/*
Description: 
a demo test of GetBestBlockHash, Get the current best block hash.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBestBlockHashTest() {
	hash, err := cli.GetBestBlockHash()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBestBlockHash", "hash", hash)
}

/*
Description:
a demo test of GetBlockHash, Get the height block hash of current best chain.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHashTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
}

/*
Description:
a demo test of GetBlockHeader, Get the special blockheader data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHeaderBytesTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	header, err := cli.GetBlockHeaderBytes(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockHeader", "header", string(*header))
}

/*
Description:
a demo test of GetBlockHeader, Get the special blockheader data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHeaderTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	header, err := cli.GetBlockHeader(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockHeader", "header", header)
}

/*
Description:
a demo test of GetBlockChainInfo, Get the blockchain state info.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockChainInfoTest() {
	info, err := cli.GetBlockChainInfo()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockChainInfo", "info", info)
}

/*
Description:
a demo test of GetBlock, Get the special block data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockBytesTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlockBytes(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockBytes", "block", string(*block))
}

/*
Description:
a demo test of GetBlock, Get the special block data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlock(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlock", "block", block)
}

/*
Description:
a demo test of GetBlockVerbose, Get the special block data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockVerboseTXTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlockVerboseTX(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockVerboseTX", "block", block)
}

/*
Description:
a demo test of GetBlockStats.
 * Author: architect.bian
 * Date: 2018/10/10 10:45
 */
func GetBlockStatsTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	result, err := cli.GetBlockStats(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockStats", "result", result)
	result4Height, err := cli.GetBlockStats(1)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockStats", "result", result4Height)
}

/*
Description:
a demo test of GetChainTips, Return information about all known tips in the block tree,
including the main chain as well as orphaned branches.
 * Author: architect.bian
 * Date: 2018/10/10 10:45
 */
func GetChainTipsTest() {
	result, err := cli.GetChainTips()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetChainTips", "result", result)
}

/*
Description:
GetChainTXStatsTest a demo test of GetChainTXStats
 * Author: architect.bian
 * Date: 2018/10/11 10:17
 */
func GetChainTXStatsTest() {
	result, err := cli.GetChainTXStats()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetChainTXStats", "result", result)
}

/*
Description:
GetChainTXStatsEntireTest a demo test of GetChainTXStatsEntire
 * Author: architect.bian
 * Date: 2018/10/11 10:18
 */
func GetChainTXStatsEntireTest() {
	hash, err := cli.GetBlockHash(105)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	result, err := cli.GetChainTXStatsEntire(100, hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetChainTXStats", "result", result)
}

/*
Description:
a demo test of GetDifficulty, Get the current difficulty for mining difficulty target.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetDifficultyTest() {
	hash, err := cli.GetDifficulty()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetDifficulty", "difficulty", hash)
}

/*
Description:
GetMempoolAncestorsTest a demo test of GetMempoolAncestors
 * Author: architect.bian
 * Date: 2018/10/11 11:40
 */
func GetMempoolAncestorsTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		result, err := cli.GetMempoolAncestors(*h)
		if err != nil {
			log.Error("", "error", err)
			panic(err)
		}
		log.Info("GetMempoolAncestors", "result", result)
		break
	}
}

/*
Description:
GetMempoolAncestorsVerboseTest a demo test of GetMempoolAncestorsVerbose
 * Author: architect.bian
 * Date: 2018/10/11 11:41
 */
func GetMempoolAncestorsVerboseTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		result, err := cli.GetMempoolAncestorsVerbose(*h)
		if err != nil {
			log.Error("", "error", err)
			panic(err)
		}
		log.Info("GetMempoolAncestorsVerbose", "result", result)
		break
	}
}

/*
Description:
GetMempoolDescendantsTest a demo test of GetMempoolDescendants
 * Author: architect.bian
 * Date: 2018/10/11 11:44
 */
func GetMempoolDescendantsTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		result, err := cli.GetMempoolDescendants(*h)
		if err != nil {
			log.Error("", "error", err)
			panic(err)
		}
		log.Info("GetMempoolDescendants", "result", result)
		break
	}
}

/*
Description:
GetMempoolDescendantsVerboseTest a demo test of GetMempoolDescendantsVerbose
 * Author: architect.bian
 * Date: 2018/10/11 11:45
 */
func GetMempoolDescendantsVerboseTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		result, err := cli.GetMempoolDescendantsVerbose(*h)
		if err != nil {
			log.Error("", "error", err)
			panic(err)
		}
		log.Info("GetMempoolDescendants", "result", result)
		break
	}
}

/*
Description:
a demo test of GetRawMempool, Get all tx ids in memory pool as a string array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetMempoolEntryTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		log.Info("GetRawMempool", "hashs", h)
		result, _ := cli.GetMempoolEntry(h.String())
		log.Info("GetMempoolEntry", "entry", result)
		break
	}
}

/*
Description: 
GetMempoolInfoTest a demo test of GetMempoolInfo
 * Author: architect.bian
 * Date: 2018/10/11 11:51
 */
func GetMempoolInfoTest() {
	result, err := cli.GetMempoolInfo()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetMempoolInfo", "result", result)
}

/*
Description:
a demo test of GetRawMempool, Get all tx ids in memory pool as a string array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetRawMempoolTest() {
	hashs, err := cli.GetRawMempool()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, h := range hashs  {
		log.Info("GetRawMempool", "hashs", h)
	}
}

/*
Description:
a demo test of GetRawMempoolVerbose, Get all tx ids in memory pool as a tx array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetRawMempoolVerboseTest() {
	result, err := cli.GetRawMempoolVerbose()
	if err != nil {
		log.Fatal("", "error", err)
	}
	for _, r := range result  {
		log.Info("GetRawMempoolVerbose", "hashs", r)
	}
}

/*
Description:
a demo test of GetTxOut, Get all tx ids in memory pool as a string array
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetTXOutTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlock(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlock", "block", block)
	for _, h := range block.Tx  {
		log.Info("block.Tx", "hashs", h)
		hs, e := results.NewHashFromStr(h)
		if e != nil {
			log.Error("", "error", e)
			panic(e)
		}
		result, _ := cli.GetTxOut(hs, 0, true)
		log.Info("GetTxOut", "result", result) //TODO bestblock value ?
		break
	}
}

/*
Description:
GetTXOutProofTest a demo test of GetTXOutProof
 * Author: architect.bian
 * Date: 2018/10/11 12:25
 */
func GetTXOutProofTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	//hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	block, err := cli.GetBlock(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlock", "block", block)
	for _, h := range block.Tx  {
		log.Info("block.Tx", "hashs", h)
		result, err := cli.GetTXOutProof(h)
		if err != nil {
			log.Error("", "error", err)
			panic(err)
		}
		log.Info("GetTXOutProof", "result", string(*result))
		break
	}
}

/*
Description:
GetTXOutSetInfoTest a demo test of GetTXOutSetInfo
 * Author: architect.bian
 * Date: 2018/10/11 12:25
 */
func GetTXOutSetInfoTest() {
	result, err := cli.GetTXOutSetInfo()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetTXOutSet", "result", result)
}

/*
Description:
PreciousBlockTest a demo test of PreciousBlock
 * Author: architect.bian
 * Date: 2018/10/11 12:31
 */
func PreciousBlockTest() {
	hash, err := cli.GetBlockHash(100)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	err2 := cli.PreciousBlock(*hash)
	if err2 != nil {
		log.Error("", "error", err2)
		panic(err2)
	}
	log.Info("PreciousBlock", "result", "success")
}

/*
Description:
PruneBlockchainTest a demo test of PruneBlockchain
 * Author: architect.bian
 * Date: 2018/10/11 12:31
 */
func PruneBlockchainTest() {
	result, err := cli.PruneBlockchain(101)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("PruneBlockchain", "result", result)
}

/*
Description:
SaveMempoolTest a demo test of SaveMempool
 * Author: architect.bian
 * Date: 2018/10/11 12:32
 */
func SaveMempoolTest() {
	err := cli.SaveMempool()
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("SaveMempool", "result", "success")
}

/*
Description:
ScanTXOutSetTest a demo test of ScanTXOutSet, not supported until now.
 * Author: architect.bian
 * Date: 2018/10/11 12:35
 */
func ScanTXOutSetTest() {
	//EXPERIMENTAL warning: this call may be removed or changed in future releases.
}

/*
Description:
a demo test of VerifyChain, Verifies blockchain database.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func VerifyChainTest() {
	verify, err := cli.VerifyChain()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("VerifyChain", "verify", verify)
}

/*
Description:
VerifyTXOutProofTest a demo test of VerifyTXOutProof
 * Author: architect.bian
 * Date: 2018/10/11 12:38
 */
func VerifyTXOutProofTest() {
	hash, err := cli.GetBlockHash(1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockHash", "hash", hash)
	block, err := cli.GetBlock(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlock", "block", block)
	for _, h := range block.Tx  {
		log.Info("block.Tx", "hashs", h)
		result, err := cli.GetTXOutProof(h)
		if err != nil {
			log.Error("", "error", err)
			panic(err)
		}
		log.Info("GetTXOutProof", "result", string(*result))
		verify, err := cli.VerifyTXOutProof(*result)
		if err != nil {
			log.Fatal("", "error", err)
		}
		log.Info("VerifyTXOutProof", "verify", verify)
		break
	}
}

/*
Description:
a demo test of VerifyChainLevel, Verifies blockchain database.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func VerifyChainLevelTest() {
	verify, err := cli.VerifyChainLevel(3)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("VerifyChain", "verify", verify)
}

/*
Description:
a demo test of VerifyChainBlocks, Verifies blockchain database.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func VerifyChainBlocksTest() {
	verify, err := cli.VerifyChainBlocks(3, 10)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("VerifyChain", "verify", verify)
}
