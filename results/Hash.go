package results

import (
	"encoding/hex"
	"fmt"
	"crypto/sha256"
)

/*
Description:
HashSize of array used to store hashes.  See Hash.
 * Author: architect.bian
 * Date: 2018/09/14 15:23
 */
const HashSize = 32

/*
Description:
MaxHashStringSize is the maximum length of a Hash hash string.
1byte=8*1bit
16 = 4*1bit
the max of 1byte is 16*16, 1byte can present with 2 hexadecimal
 * Author: architect.bian
 * Date: 2018/09/14 15:23
 */
const MaxHashStringSize = HashSize * 2

/*
Description:
HashStrSizeErr describes an error that indicates the caller specified a hash
string that has too many characters.
 * Author: architect.bian
 * Date: 2018/09/14 15:29
 */
var HashStrSizeErr = fmt.Errorf("max hash string length is %v bytes", MaxHashStringSize)

/*
Description:
Hash is used in several of the bitcoin messages and common structures.  It
typically represents the double sha256 of data.
reference document https://bitcoin.org/en/developer-reference#hash-byte-order
 * Author: architect.bian
 * Date: 2018/08/27 11:10
 */
type Hash [HashSize]byte

/*
Description:
String returns the Hash as the hexadecimal string of the byte-reversed hash.
 * Author: architect.bian
 * Date: 2018/08/27 11:15
 */
func (hash Hash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

/*
Description:
CloneBytes returns a copy of the bytes which represent the hash as a byte slice.

NOTE: It is generally cheaper to just slice the hash directly thereby reusing
the same bytes rather than calling this method.
 * Author: architect.bian
 * Date: 2018/08/27 11:16
 */
func (hash *Hash) CloneBytes() []byte {
	newHash := make([]byte, HashSize)
	copy(newHash, hash[:])
	return newHash
}

/*
Description:
SetBytes sets the bytes which represent the hash.  An error is returned if
the number of bytes passed in is not HashSize.
 * Author: architect.bian
 * Date: 2018/08/27 11:18
 */
func (hash *Hash) SetBytes(newHash []byte) error {
	nhlen := len(newHash)
	if nhlen != HashSize {
		return fmt.Errorf("invalid hash length of %v, want %v", nhlen, HashSize)
	}
	copy(hash[:], newHash)
	return nil
}

/*
Description:
IsEqual returns true if target is the same as hash.
 * Author: architect.bian
 * Date: 2018/08/27 11:19
 */
func (hash *Hash) IsEqual(target *Hash) bool {
	if hash == nil && target == nil {
		return true
	}
	if hash == nil || target == nil {
		return false
	}
	return *hash == *target
}

/*
Description:
NewHash returns a new Hash from a byte slice.  An error is returned if
the number of bytes passed in is not HashSize.
 * Author: architect.bian
 * Date: 2018/08/27 11:20
 */
func NewHash(newHash []byte) (*Hash, error) {
	var sh Hash
	err := sh.SetBytes(newHash)
	if err != nil {
		return nil, err
	}
	return &sh, err
}

/*
Description:
NewHashFromStr creates a Hash from a hash string.  The string should be
the hexadecimal string of a byte-reversed hash, but any missing characters
result in zero padding at the end of the Hash.
 * Author: architect.bian
 * Date: 2018/08/27 11:30
 */
func NewHashFromStr(hash string) (*Hash, error) {
	ret := new(Hash)
	err := Decode(ret, hash)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

/*
Description:
Decode decodes the byte-reversed hexadecimal string encoding of a Hash to a destination.
 * Author: architect.bian
 * Date: 2018/08/27 11:30
 */
func Decode(dst *Hash, src string) error {
	// Return error if hash string is too long.
	if len(src) > MaxHashStringSize {
		return HashStrSizeErr
	}

	// Hex decoder expects the hash to be a multiple of two.  When not, pad with a leading zero.
	var srcBytes []byte
	if len(src)%2 == 0 {
		srcBytes = []byte(src)
	} else {
		srcBytes = make([]byte, 1+len(src))
		srcBytes[0] = '0'
		copy(srcBytes[1:], src)
	}

	// Hex decode the source bytes to a temporary destination.
	var reversedHash Hash
	_, err := hex.Decode(reversedHash[HashSize-hex.DecodedLen(len(srcBytes)):], srcBytes)
	if err != nil {
		return err
	}

	// Reverse copy from the temporary hash to destination.  Because the
	// temporary was zeroed, the written result will be correctly padded.
	for i, b := range reversedHash[:HashSize/2] {
		dst[i], dst[HashSize-1-i] = reversedHash[HashSize-1-i], b
	}

	return nil
}

/*
Description:
HashB calculates hash(b) and returns the resulting bytes.
 * Author: architect.bian
 * Date: 2018/08/27 13:05
 */
func HashB(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

/*
Description:
HashH calculates hash(b) and returns the resulting bytes as a Hash.
 * Author: architect.bian
 * Date: 2018/08/27 13:05
 */
func HashH(b []byte) Hash {
	return Hash(sha256.Sum256(b))
}

/*
Description:
DoubleHashB calculates hash(hash(b)) and returns the resulting bytes.
 * Author: architect.bian
 * Date: 2018/08/27 13:05
 */
func DoubleHashB(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

/*
Description:
DoubleHashH calculates hash(hash(b)) and returns the resulting bytes as a Hash.
 * Author: architect.bian
 * Date: 2018/08/27 13:05
 */
func DoubleHashH(b []byte) Hash {
	first := sha256.Sum256(b)
	return Hash(sha256.Sum256(first[:]))
}
