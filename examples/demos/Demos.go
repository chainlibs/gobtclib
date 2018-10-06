package demos

import (
	"github.com/chainlibs/gobtclib/client"
)

var cli *client.Client

/*
Description:
Initialize create an instance of client.Client with client.Config
 * Author: architect.bian
 * Date: 2018/09/02 18:24
 */
func Initialize(cfg *client.Config) {
	cli = client.NewClient(cfg).Startup()
}

/*
Description:
shutdown client
 * Author: architect.bian
 * Date: 2018/09/14 12:07
 */
func Shutdown() {
	cli.Shutdown()
}
