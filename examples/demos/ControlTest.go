package demos

import (
	"github.com/gobasis/log"
	//"github.com/chainlibs/gobtclib/results"
)

/*
Description:
a demo test of GetMemoryInfo, Get the memory info from server peer.
 * Author: architect.bian
 * Date: 2018/10/15 11:19
 */
func GetMemoryInfoTest() {
	result, err := cli.GetMemoryInfo()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetMemoryInfo", "result", result)
}

/*
Description: 
GetMemoryInfo4MallocInfoTest a demo test of GetMemoryInfo with argument:
model as mallocinfo
 * Author: architect.bian
 * Date: 2018/10/15 11:08
 */
func GetMemoryInfo4MallocInfoTest() {
	result, err := cli.GetMemoryInfo4MallocInfo()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("GetMemoryInfo", "result", result)
}

/*
Description:
a demo test of Help.
 * Author: architect.bian
 * Date: 2018/10/15 11:18
 */
func HelpTest() {
	result, err := cli.Help()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("Help", "result", result)
}

///*
//Description:
//a demo test of Logging.
// * Author: architect.bian
// * Date: 2018/09/02 18:24
// */
//func LoggingTest() {
//	result, err := cli.Logging()
//	if err != nil {
//		log.Fatal("", "error", err)
//	}
//	log.Info("Logging", "result", result)
//}

/*
Description:
a demo test of Stop.
 * Author: architect.bian
 * Date: 2018/10/15 11:18
 */
func StopTest() {
	err := cli.Stop()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("Stop", "result", "success")
}

/*
Description:
a demo test of Uptime.
 * Author: architect.bian
 * Date: 2018/10/15 11:17
 */
func UptimeTest() {
	result, err := cli.Uptime()
	if err != nil {
		log.Fatal("", "error", err)
	}
	log.Info("Uptime", "result", result)
}
