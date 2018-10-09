package main

import (
	"github.com/chainlibs/gobtclib/client"
	"github.com/chainlibs/gobtclib/examples/demos"
	"github.com/gobasis/log"
	"github.com/gobasis/log/zapimpl"
)

/*
Description:
show demos of gobtclib/client
 * Author: architect.bian
 * Date: 2018/09/02 17:40
 */
func main() {
	log.UseLog(&zapimpl.Logger{}) // use zap log
	log.SetLevel(log.DevDebugLevel)
	log.Info("start up bitcoin rpc client")
	//cfg := &client.Config{
	//	Host:         "172.16.2.35:3333",
	//	User:         "et",
	//	Pass:         "www.et.com",
	//}
	cfg := &client.Config{
		Host:         "172.16.2.27:8332",
		User:         "btc",
		Pass:         "btcpwd",
	}
	demos.Initialize(cfg)
	defer demos.Shutdown()
	//demos.GetBestBlockHashTest()
	//demos.GetBlockBytesTest()
	//demos.GetBlockTest()
	//demos.GetBlockVerboseTXTest()
	//demos.GetBlockVerboseTXTest()
	//demos.GetBlockChainInfoTest()
	//demos.GetBlockCountTest()
	//demos.GetBlockHashTest()
	//demos.GetBlockBytesTest()
	demos.GetBlockHeaderTest()
	//getchaintips //TODO
	//getchaintxstats
	//demos.GetDifficultyTest()
	//demos.RescanBlocksTest() //TODO failed
	//demos.GetCFilter() //TODO unimplemented
	//demos.InvalidateBlockTest() //very dangerous
	//demos.VerifyChainTest()
	//demos.VerifyChainLevelTest()
	//demos.VerifyChainBlocksTest()
	//demos.GetRawMempoolTest()
	//demos.GetRawMempoolVerboseTest()
	//demos.GetMempoolEntryTest()
	//demos.GetTxOutTest()
}