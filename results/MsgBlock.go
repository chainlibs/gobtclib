package results

import "io"

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
