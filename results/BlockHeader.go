package results

import (
	"time"
	"io"
)

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

