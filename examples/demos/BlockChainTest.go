package demos

import (
	"github.com/gobasis/log"
)

/*
Description:
a demo test of GetBlockCount
 * Author: architect.bian
 * Date: 2018/09/02 18:24
 */
func GetBlockCountTest() {
	// Get the current block count.
	blockCount, err := cli.GetBlockCount()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetBlockCount", "count", blockCount)
}