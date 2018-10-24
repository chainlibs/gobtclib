package btc

import "github.com/chainlibs/gobtclib/client"

/*
Description:
global client instance
 * Author: architect.bian
 * Date: 2018/10/24 12:59
 */
var Client *client.Client

/*
Description:
initialize a Client instance with host/user/passwd parameters
 * Author: architect.bian
 * Date: 2018/10/24 12:57
 */
func StartUp(host, user, passwd string) {
	cfg := &client.Config{
		Host:         host,
		User:         user,
		Pass:         passwd,
	}
	Client = client.NewClient(cfg).Startup()
}