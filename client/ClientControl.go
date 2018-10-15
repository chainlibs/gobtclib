package client

import (
	"github.com/chainlibs/gobtclib/futures"
)

/*
Description:
GetMemoryInfoAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetMemoryInfo for more details.
 * Author: architect.bian
 * Date: 2018/09/14 13:18
 */
func (c *Client) GetMemoryInfoAsync(model string) futures.FutureResult {
	cmd := NewCommand("getmemoryinfo", model)
	return c.sendCmd(cmd)
}

/*
Description:
GetMemoryInfo returns general statistics about memory usage in the daemon.
 * Author: architect.bian
 * Date: 2018/09/14 13:15
 */
func (c *Client) GetMemoryInfo() (*interface{}, error) {
	return c.GetMemoryInfoAsync("stats").Receive()
}

/*
Description:
GetMemoryInfo4MallocInfo returns an XML string describing
low-level heap state (only available if compiled with glibc 2.10+)
 * Author: architect.bian
 * Date: 2018/09/14 13:15
 */
func (c *Client) GetMemoryInfo4MallocInfo() (*string, error) {
	return futures.FutureString(c.GetMemoryInfoAsync("mallocinfo")).Receive()
}

/*
Description:
HelpAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See Help for more details.
 * Author: architect.bian
 * Date: 2018/10/15 10:40
 */
func (c *Client) HelpAsync() futures.FutureString {
	cmd := NewCommand("help")
	return c.sendCmd(cmd)
}

/*
Description:
Help List all commands, or get help for a specified command.
 * Author: architect.bian
 * Date: 2018/10/15 10:40
 */
func (c *Client) Help() (*string, error) {
	return c.HelpAsync().Receive()
}

/*
Description:
StopAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See Stop for more details.
 * Author: architect.bian
 * Date: 2018/10/15 10:38
 */
func (c *Client) StopAsync() futures.FutureResult {
	cmd := NewCommand("stop")
	return c.sendCmd(cmd)
}

/*
Description:
Stop Bitcoin server.
 * Author: architect.bian
 * Date: 2018/10/15 10:38
 */
func (c *Client) Stop() (error) {
	_, err := c.StopAsync().Receive()
	if err != nil {
		return err
	}
	return nil
}

/*
Description:
GetUptimeAsync returns an instance of a type that can be used to get
the result of the RPC at some future time by invoking the Receive function on
the returned instance.

See GetUptime for more details.
 * Author: architect.bian
 * Date: 2018/10/15 10:36
 */
func (c *Client) UptimeAsync() futures.FutureInt32 {
	cmd := NewCommand("uptime")
	return c.sendCmd(cmd)
}

/*
Description:
Uptime Returns the total uptime of the server.
 * Author: architect.bian
 * Date: 2018/10/15 10:36
 */
func (c *Client) Uptime() (int32, error) {
	return c.UptimeAsync().Receive()
}
