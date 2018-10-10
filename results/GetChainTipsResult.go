package results

/*
Description: 
GetChainTipsResult models the data from the getchaintips command
 * Author: architect.bian
 * Date: 2018/10/10 11:18
 */
type GetChainTipsResult struct {
	Height    int32  `json:"height"`
	Hash      string `json:"hash"`
	Branchlen int32  `json:"branchlen"`
	Status    string `json:"status"`
}
