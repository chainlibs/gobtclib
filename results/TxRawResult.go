package results

import (
	"encoding/json"
)

// Vin models parts of the tx data.  It is defined separately since
// getrawtransaction, decoderawtransaction, and searchrawtransaction use the
// same structure.
type Vin struct {
	Coinbase  string     `json:"coinbase"`
	Txid      string     `json:"txid"`
	Vout      uint32     `json:"vout"`
	ScriptSig *ScriptSig `json:"scriptSig"`
	Sequence  uint32     `json:"sequence"`
	Witness   []string   `json:"txinwitness"`
}

// IsCoinBase returns a bool to show if a Vin is a Coinbase one or not.
func (v *Vin) IsCoinBase() bool {
	return len(v.Coinbase) > 0
}

// HasWitness returns a bool to show if a Vin has any witness data associated
// with it or not.
func (v *Vin) HasWitness() bool {
	return len(v.Witness) > 0
}

// ScriptSig models a signature script.  It is defined separately since it only
// applies to non-coinbase.  Therefore the field in the Vin structure needs
// to be a pointer.
type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

// MarshalJSON provides a custom Marshal method for Vin.
func (v *Vin) MarshalJSON() ([]byte, error) {
	if v.IsCoinBase() {
		coinbaseStruct := struct {
			Coinbase string   `json:"coinbase"`
			Sequence uint32   `json:"sequence"`
			Witness  []string `json:"witness,omitempty"`
		}{
			Coinbase: v.Coinbase,
			Sequence: v.Sequence,
			Witness:  v.Witness,
		}
		return json.Marshal(coinbaseStruct)
	}

	if v.HasWitness() {
		txStruct := struct {
			Txid      string     `json:"txid"`
			Vout      uint32     `json:"vout"`
			ScriptSig *ScriptSig `json:"scriptSig"`
			Witness   []string   `json:"txinwitness"`
			Sequence  uint32     `json:"sequence"`
		}{
			Txid:      v.Txid,
			Vout:      v.Vout,
			ScriptSig: v.ScriptSig,
			Witness:   v.Witness,
			Sequence:  v.Sequence,
		}
		return json.Marshal(txStruct)
	}

	txStruct := struct {
		Txid      string     `json:"txid"`
		Vout      uint32     `json:"vout"`
		ScriptSig *ScriptSig `json:"scriptSig"`
		Sequence  uint32     `json:"sequence"`
	}{
		Txid:      v.Txid,
		Vout:      v.Vout,
		ScriptSig: v.ScriptSig,
		Sequence:  v.Sequence,
	}
	return json.Marshal(txStruct)
}

// Vout models parts of the tx data.  It is defined separately since both
// getrawtransaction and decoderawtransaction use the same structure.
type Vout struct {
	Value        float64            `json:"value"`
	N            uint32             `json:"n"`
	ScriptPubKey ScriptPubKeyResult `json:"scriptPubKey"`
}

// ScriptPubKeyResult models the scriptPubKey data of a tx script.  It is
// defined separately since it is used by multiple commands.
type ScriptPubKeyResult struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex,omitempty"`
	ReqSigs   int32    `json:"reqSigs,omitempty"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses,omitempty"`
}

// TxRawResult models the data from the getrawtransaction command.
type TxRawResult struct {
	Hex           string `json:"hex"`
	Txid          string `json:"txid"`
	Hash          string `json:"hash,omitempty"`
	Size          int32  `json:"size,omitempty"`
	Vsize         int32  `json:"vsize,omitempty"`
	Version       int32  `json:"version"`
	LockTime      uint32 `json:"locktime"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	BlockHash     string `json:"blockhash,omitempty"`
	Confirmations uint64 `json:"confirmations,omitempty"`
	Time          int64  `json:"time,omitempty"`
	Blocktime     int64  `json:"blocktime,omitempty"`
}

