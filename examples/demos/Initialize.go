package demos

import (
	"github.com/chainlibs/gobtclib/client"
)

var cli *client.Client

/*
Description:
initialize an instance of client.Client with client.Config
 * Author: architect.bian
 * Date: 2018/09/02 18:24
 */
func Initialize(cfg *client.Config) {
	cli = client.NewClient(cfg).Startup()
	//defer cli.Shutdown()
}
