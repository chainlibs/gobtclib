package demos

import (
	"github.com/gobasis/log"
)

/*
Description:
A demo test of rpc method generate with method Send.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GenerateTest() {
	result, err := cli.Send("generate", 1)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("generate", "result", result)
}

/*
Description:
A demo test of rpc method generatetoaddress with method Send.
* Author: architect.bian
* Date: 2018/10/15 12:14
*/
func GenerateToAddressTest() {
	result, err := cli.Send("generatetoaddress", 1, addressTo)
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("generatetoaddress", "result", result)
}
