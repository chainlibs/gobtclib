package results

/*
Description:
GetMempoolInfoResult models the data from the GetMempoolInfo command
 * Author: architect.bian
 * Date: 2018/10/11 11:37
 */
type GetMempoolInfoResult struct {
	Size          int32		`json:"size"`
	Bytes         int32		`json:"bytes"`
	Usage         int32		`json:"usage"`
	MaxMempool    int32		`json:"maxmempool"`
	MempoolMinfee float32	`json:"mempoolminfee"`
	MinRelayTXFee float32	`json:"minrelaytxfee"`
}
