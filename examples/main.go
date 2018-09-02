package main

import (
	"fmt"
	"github.com/chainlibs/gobtclib/client"
	"github.com/chainlibs/gobtclib/examples/demos"
)

/*
Description:
show demos of gobtclib/client
 * Author: architect.bian
 * Date: 2018/09/02 17:40
 */
func main() {
	fmt.Println("start up bitcoin rpc client")
	cfg := &client.Config{
		Host:         "172.16.2.41:8332",
		User:         "user",
		Pass:         "password",
	}
	demos.Initialize(cfg)
	demos.GetBlockCountTest()
}