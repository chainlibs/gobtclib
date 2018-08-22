package clients

import "fmt"

// BitcoinNet represents which bitcoin network a message belongs to.
type BitcoinNet uint32

// Constants used to indicate the message bitcoin network.  They can also be
// used to seek to the next message when a stream's state is unknown, but
// this package does not provide that functionality since it's generally a
// better idea to simply disconnect clients that are misbehaving over TCP.
const (
	// MainNet represents the main bitcoin network.
	MainNet BitcoinNet = 0xd9b4bef9

	// TestNet represents the regression test network.
	TestNet BitcoinNet = 0xdab5bffa

	// TestNet3 represents the test network (version 3).
	TestNet3 BitcoinNet = 0x0709110b

	// SimNet represents the simulation test network.
	SimNet BitcoinNet = 0x12141c16
)

// bnStrings is a map of bitcoin networks back to their constant names for
// pretty printing.
var bnStrings = map[BitcoinNet]string{
	MainNet:  "MainNet",
	TestNet:  "TestNet",
	TestNet3: "TestNet3",
	SimNet:   "SimNet",
}

// String returns the BitcoinNet in human-readable form.
func (n BitcoinNet) String() string {
	if s, ok := bnStrings[n]; ok {
		return s
	}

	return fmt.Sprintf("Unknown BitcoinNet (%d)", uint32(n))
}

