package demos

import (
	"github.com/gobasis/log"
)

/*
Description:
A demo test of rpc method getnetworkinfo with method Send.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GetNetworkInfoTest() {
	result, err := cli.Send("getnetworkinfo")
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("getnetworkinfo", "result", result)
}
