package results

/*
Description:
GetTXOutSetInfoResult models the data from the GetTXOutSetInfo command
 * Author: architect.bian
 * Date: 2018/10/11 11:36
*/
type GetTXOutSetInfoResult struct {
	Height          int32   `json:"height"`
	Bestblock       string    `json:"bestblock"`
	Transactions    int32   `json:"transactions"`
	Txouts          int32   `json:"txouts"`
	Bogosize        int32   `json:"bogosize"`
	HashSerialized2 string  `json:"hash_serialized_2"`
	DiskSize        int32   `json:"disk_size"`
	TotalAmount     float32 `json:"total_amount"`
}
