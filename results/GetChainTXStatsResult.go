package results

/*
Description:
GetChainTXStatsResult models the data from the GetChainTXStats command
 * Author: architect.bian
 * Date: 2018/10/11 11:36
 */
type GetChainTXStatsResult struct {
	Time                 int32   `json:"time"`
	Txcount              int32   `json:"txcount"`
	WindowFinalBlockHash string  `json:"window_final_block_hash"`
	WindowBlockCount     int32   `json:"window_block_count"`
	WindowTXCount        int32   `json:"window_tx_count"`
	WindowInterval       int32   `json:"window_interval"`
	Txrate               float32 `json:"txrate"`
}
