package results

/*
Description: 
GetBlockStatsResult models the data from the GetBlockStats command
 * Author: architect.bian
 * Date: 2018/10/11 11:36
 */
type GetBlockStatsResult struct {
	Avgfee              int32		`json:"avgfee,omitempty"`
	Avgfeerate          int32		`json:"avgfeerate"`
	Avgtxsize           int32		`json:"avgtxsize"`
	Blockhash           string		`json:"blockhash"`
	FeeratePercentiles []int32		`json:"feerate_percentiles"`
	Height              int32		`json:"height"`
	Ins                 int32		`json:"ins"`
	Maxfee              int32		`json:"maxfee"`
	Maxfeerate          int32		`json:"maxfeerate"`
	Maxtxsize           int32		`json:"maxtxsize"`
	Medianfee           int32		`json:"medianfee"`
	Mediantime          int32		`json:"mediantime"`
	Mediantxsize        int32		`json:"mediantxsize"`
	Minfee              int32		`json:"minfee"`
	Minfeerate          int32		`json:"minfeerate"`
	Mintxsize           int32		`json:"mintxsize"`
	Outs                int32		`json:"outs"`
	Subsidy             int32		`json:"subsidy"`
	SwtotalSize        int32		`json:"swtotal_size"`
	SwtotalWeight      int32		`json:"swtotal_weight"`
	Swtxs               int32		`json:"swtxs"`
	Time                int32		`json:"time"`
	TotalOut           int32		`json:"total_out"`
	TotalSize          int32		`json:"total_size"`
	TotalWeight        int32		`json:"total_weight"`
	Totalfee            int32		`json:"totalfee"`
	Txs                 int32		`json:"txs"`
	UtxoIncrease       int32		`json:"utxo_increase"`
	UtxoSizeInc       int32		`json:"utxo_size_inc"`
}
