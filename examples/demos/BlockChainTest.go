package demos

import (
	"github.com/gobasis/log"
	"github.com/chainlibs/gobtclib/client"
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
a demo test of GetBlockHeaderVerbose, Get the special blockheader data structure.
 * Author: architect.bian
 * Date: 2018/09/14 13:13
 */
func GetBlockHeaderVerboseTest() {
	//hash, err := cli.GetBlockHash(1)
	//if err != nil {
	//	log.Fatal("", "error", err)
	//}
	//log.Info("GetBlockHash", "hash", hash)
	hash, _ := client.NewHashFromStr("353eb168c70770b281f634fd0d22fbaee94b3fed1d76276e91165444e876f658")
	header, err := cli.GetBlockHeaderVerbose(hash)
	if err != nil {
		log.Error("", "error", err)
		panic(err)
	}
	log.Info("GetBlockHeaderVerbose", "header", header)
}