package results

/*
Description:
SoftForkDescription describes the current state of a soft-fork which was
deployed using a super-majority block signalling.
 * Author: architect.bian
 * Date: 2018/10/07 17:49
 */
type SoftForkDescription struct {
	ID      string `json:"id"`
	Version uint32 `json:"version"`
	Reject  struct {
		Status bool `json:"status"`
	} `json:"reject"`
}

/*
Description:
Bip9SoftForkDescription describes the current state of a defined BIP0009
version bits soft-fork.
 * Author: architect.bian
 * Date: 2018/10/07 17:49
 */
type Bip9SoftForkDescription struct {
	Status    string `json:"status"`
	Bit       uint8  `json:"bit"`
	StartTime int64  `json:"startTime"`
	Timeout   int64  `json:"timeout"`
	Since     int32  `json:"since"`
}

/*
Description:
GetBlockChainInfoResult models the data returned from the getblockchaininfo command.
 * Author: architect.bian
 * Date: 2018/09/17 14:53
 */
type GetBlockChainInfoResult struct {
	Chain                string                              `json:"chain"`
	Blocks               int32                               `json:"blocks"`
	Headers              int32                               `json:"headers"`
	BestBlockHash        string                              `json:"bestblockhash"`
	Difficulty           float64                             `json:"difficulty"`
	MedianTime           int64                               `json:"mediantime"`
	VerificationProgress float64                             `json:"verificationprogress,omitempty"`
	Pruned               bool                                `json:"pruned"`
	PruneHeight          int32                               `json:"pruneheight,omitempty"`
	ChainWork            string                              `json:"chainwork,omitempty"`
	SoftForks            []*SoftForkDescription              `json:"softforks"`
	Bip9SoftForks        map[string]*Bip9SoftForkDescription `json:"bip9_softforks"`
}

