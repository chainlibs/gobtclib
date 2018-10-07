package results

/*
Description:
GetBlockHeaderVerboseResult models the data from the getblockheader command when
the verbose flag is set.  When the verbose flag is not set, getblockheader
returns a hex-encoded string.
 * Author: architect.bian
 * Date: 2018/10/07 17:37
 */
type GetBlockHeaderVerboseResult struct {
	Hash          string  `json:"hash"`
	Confirmations uint64  `json:"confirmations"`
	Height        int32   `json:"height"`
	Version       int32   `json:"version"`
	VersionHex    string  `json:"versionHex"`
	MerkleRoot    string  `json:"merkleroot"`
	Time          int64   `json:"time"`
	Nonce         uint64  `json:"nonce"`
	Bits          string  `json:"bits"`
	Difficulty    float64 `json:"difficulty"`
	PreviousHash  string  `json:"previousblockhash,omitempty"`
	NextHash      string  `json:"nextblockhash,omitempty"`
}

