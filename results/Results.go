package results

import (
	"encoding/json"
	"time"
	"io"
)

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

/*
Description:
BlockHeader defines information about a block and is used in the bitcoin
block (MsgBlock) and headers (MsgHeaders) messages.
 * Author: architect.bian
 * Date: 2018/09/17 15:27
 */
type BlockHeader struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32

	// Hash of the previous block header in the block chain.
	PrevBlock Hash

	// Merkle tree reference to hash of all transactions for the block.
	MerkleRoot Hash

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	Timestamp time.Time

	// Difficulty target for the block.
	Bits uint32

	// Nonce used to generate the block.
	Nonce uint32
}

// Deserialize decodes a block header from r into the receiver using a format
// that is suitable for long-term storage such as a database while respecting
// the Version field.
func (h *BlockHeader) Deserialize(r io.Reader) error { //TODO parse blockheader data
	// At the current time, there is no difference between the wire encoding
	// at protocol version 0 and the stable long-term storage format.  As
	// a result, make use of readBlockHeader.
	return nil
}

// Serialize encodes a block header from r into the receiver using a format
// that is suitable for long-term storage such as a database while respecting
// the Version field.
func (h *BlockHeader) Serialize(w io.Writer) error {
	// At the current time, there is no difference between the wire encoding
	// at protocol version 0 and the stable long-term storage format.  As
	// a result, make use of writeBlockHeader.
	return nil
}

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

// TxIn defines a bitcoin transaction input.
type TxIn struct {
	PreviousOutPoint OutPoint
	SignatureScript  []byte
	Witness          TxWitness
	Sequence         uint32
}

// OutPoint defines a bitcoin data type that is used to track previous
// transaction outputs.
type OutPoint struct {
	Hash  Hash
	Index uint32
}

// TxWitness defines the witness for a TxIn. A witness is to be interpreted as
// a slice of byte slices, or a stack with one or many elements.
type TxWitness [][]byte

// TxOut defines a bitcoin transaction output.
type TxOut struct {
	Value    int64
	PkScript []byte
}

/*
Description:
MsgTx implements the Message interface and represents a bitcoin tx message.
It is used to deliver transaction information in response to a getdata
message (MsgGetData) for a given transaction.

Use the AddTxIn and AddTxOut functions to build up the list of transaction
inputs and outputs.
 * Author: architect.bian
 * Date: 2018/10/07 17:50
 */
type MsgTx struct {
	Version  int32
	TxIn     []*TxIn
	TxOut    []*TxOut
	LockTime uint32
}

// MsgBlock implements the Message interface and represents a bitcoin
// block message.  It is used to deliver block and transaction information in
// response to a getdata message (MsgGetData) for a given block hash.
type MsgBlock struct {
	Header       BlockHeader
	Transactions []*MsgTx
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
// See Deserialize for decoding blocks stored to disk, such as in a database, as
// opposed to decoding blocks from the wire.
func (msg *MsgBlock) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error { //TODO parse byte data
	//err := readBlockHeader(r, pver, &msg.Header)
	//if err != nil {
	//	return err
	//}
	//
	//txCount, err := ReadVarInt(r, pver)
	//if err != nil {
	//	return err
	//}
	//
	//// Prevent more transactions than could possibly fit into a block.
	//// It would be possible to cause memory exhaustion and panics without
	//// a sane upper bound on this count.
	//if txCount > maxTxPerBlock {
	//	str := fmt.Sprintf("too many transactions to fit into a block "+
	//		"[count %d, max %d]", txCount, maxTxPerBlock)
	//	return messageError("MsgBlock.BtcDecode", str)
	//}
	//
	//msg.Transactions = make([]*MsgTx, 0, txCount)
	//for i := uint64(0); i < txCount; i++ {
	//	tx := MsgTx{}
	//	err := tx.BtcDecode(r, pver, enc)
	//	if err != nil {
	//		return err
	//	}
	//	msg.Transactions = append(msg.Transactions, &tx)
	//}
	//
	return nil
}

// MessageEncoding represents the wire message encoding format to be used.
type MessageEncoding uint32

const (
	// BaseEncoding encodes all messages in the default format specified
	// for the Bitcoin wire protocol.
	BaseEncoding MessageEncoding = 1 << iota

	// WitnessEncoding encodes all messages other than transaction messages
	// using the default Bitcoin wire protocol specification. For transaction
	// messages, the new encoding format detailed in BIP0144 will be used.
	WitnessEncoding
)
// Deserialize decodes a block from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.  This function differs from BtcDecode in that
// BtcDecode decodes from the bitcoin wire protocol as it was sent across the
// network.  The wire encoding can technically differ depending on the protocol
// version and doesn't even really need to match the format of a stored block at
// all.  As of the time this comment was written, the encoded block is the same
// in both instances, but there is a distinct difference and separating the two
// allows the API to be flexible enough to deal with changes.
func (msg *MsgBlock) Deserialize(r io.Reader) error {
	// At the current time, there is no difference between the wire encoding
	// at protocol version 0 and the stable long-term storage format.  As
	// a result, make use of BtcDecode.
	//
	// Passing an encoding type of WitnessEncoding to BtcEncode for the
	// MessageEncoding parameter indicates that the transactions within the
	// block are expected to be serialized according to the new
	// serialization structure defined in BIP0141.
	return msg.BtcDecode(r, 0, WitnessEncoding)
}

// DeserializeNoWitness decodes a block from r into the receiver similar to
// Deserialize, however DeserializeWitness strips all (if any) witness data
// from the transactions within the block before encoding them.
func (msg *MsgBlock) DeserializeNoWitness(r io.Reader) error {
	return msg.BtcDecode(r, 0, BaseEncoding)
}

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

/*
Description:
GetBlockVerboseResult models the data from the getblock command when the
verbose flag is 0/1.  When the verbose flag is set as 0, getblock returns a
hex-encoded string.
 * Author: architect.bian
 * Date: 2018/09/17 16:27
 */
type GetBlockVerboseResult struct {
	Hash          string        `json:"hash"`
	Confirmations uint64        `json:"confirmations"`
	StrippedSize  int32         `json:"strippedsize"`
	Size          int32         `json:"size"`
	Weight        int32         `json:"weight"`
	Height        int64         `json:"height"`
	Version       int32         `json:"version"`
	VersionHex    string        `json:"versionHex"`
	MerkleRoot    string        `json:"merkleroot"`
	Tx            []string      `json:"tx,omitempty"`
	RawTx         []TxRawResult `json:"rawtx,omitempty"`
	Time          int64         `json:"time"`
	Nonce         uint32        `json:"nonce"`
	Bits          string        `json:"bits"`
	Difficulty    float64       `json:"difficulty"`
	PreviousHash  string        `json:"previousblockhash"`
	NextHash      string        `json:"nextblockhash,omitempty"`
}

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

/*
Description:
GetRawMempoolVerboseResult models the data returned from the getrawmempool
command when the verbose flag is set.  When the verbose flag is not set,
getrawmempool returns an array of transaction hashes.
 * Author: architect.bian
 * Date: 2018/09/17 20:12
 */
type GetRawMempoolVerboseResult struct {
	Size             int32    `json:"size"`
	Vsize            int32    `json:"vsize"`
	Fee              float64  `json:"fee"`
	Time             int64    `json:"time"`
	Height           int64    `json:"height"`
	StartingPriority float64  `json:"startingpriority"` //TODO new version compatibility
	CurrentPriority  float64  `json:"currentpriority"`  //TODO new version compatibility
	Depends          []string `json:"depends"`
}

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

// GetTxOutResult models the data from the gettxout command.
type GetTxOutResult struct {
	BestBlock     string             `json:"bestblock"`
	Confirmations int64              `json:"confirmations"`
	Value         float64            `json:"value"`
	ScriptPubKey  ScriptPubKeyResult `json:"scriptPubKey"`
	Coinbase      bool               `json:"coinbase"`
}

