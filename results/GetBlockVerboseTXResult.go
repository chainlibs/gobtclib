package results

/*
Description:
GetBlockVerboseTXResult models the data from the getblock command when the
verbose flag is 2.  When the verbose flag is set as 0, getblock returns a
hex-encoded string.
 * Author: architect.bian
 * Date: 2018/09/17 16:27
 */
type GetBlockVerboseTXResult struct {
	GetBlockVerboseResult
	Tx []TxRawResult 			`json:"tx"`
}

