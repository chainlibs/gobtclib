package results

/*
Description:
GetRawMempoolVerboseResult models the data returned from the getrawmempool
command when the verbose flag is set.  When the verbose flag is not set,
getrawmempool returns an array of transaction hashes.
 * Author: architect.bian
 * Date: 2018/09/17 20:12
 */
type GetRawMempoolVerboseResult struct {
	Size             int32    `json:"size"`
	Vsize            int32    `json:"vsize"`
	Fee              float64  `json:"fee"`
	Time             int64    `json:"time"`
	Height           int64    `json:"height"`
	StartingPriority float64  `json:"startingpriority"` //TODO new version compatibility
	CurrentPriority  float64  `json:"currentpriority"`  //TODO new version compatibility
	Depends          []string `json:"depends"`
}

