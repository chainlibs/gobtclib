package results

/*
Description:
GetMempoolEntryResult models the data returned from the getmempoolentry command.
 * Author: architect.bian
 * Date: 2018/09/17 20:24
 */
type GetMempoolEntryResult struct {
	Size             int32    `json:"size"`
	Fee              float64  `json:"fee"`
	ModifiedFee      float64  `json:"modifiedfee"`
	Time             int64    `json:"time"`
	Height           int64    `json:"height"`
	StartingPriority float64  `json:"startingpriority"`
	CurrentPriority  float64  `json:"currentpriority"`
	DescendantCount  int64    `json:"descendantcount"`
	DescendantSize   int64    `json:"descendantsize"`
	DescendantFees   float64  `json:"descendantfees"`
	AncestorCount    int64    `json:"ancestorcount"`
	AncestorSize     int64    `json:"ancestorsize"`
	AncestorFees     float64  `json:"ancestorfees"`
	Depends          []string `json:"depends"`
}

