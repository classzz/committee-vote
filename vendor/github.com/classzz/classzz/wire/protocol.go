// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"strconv"
	"strings"
)

// XXX pedro: we will probably need to bump this.
const (
	// ProtocolVersion is the latest protocol version this package supports.
	ProtocolVersion uint32 = 70015
)

// ServiceFlag identifies services supported by a bitcoin peer.
type ServiceFlag uint64

const (
	// SFNodeNetwork is a flag used to indicate a peer is a full node.
	SFNodeNetwork ServiceFlag = 1 << iota

	// SFNodeGetUTXO is a flag used to indicate a peer supports the
	// getutxos and utxos commands (BIP0064).
	SFNodeGetUTXO

	// SFNodeBloom is a flag used to indicate a peer supports bloom
	// filtering.
	SFNodeBloom

	// SFNodeWitness is a flag used to indicate a peer supports blocks
	// and transactions including witness data (BIP0144).
	SFNodeWitness

	// SFNodeXthin is a flag used to indicate a peer supports xthin blocks.
	SFNodeXthin

	// SFNodeBitcoinCash indicates a node is running on the Bitcoin Cash
	// network. Bitcoin Core peers should disconnect upon seeing this service bit.
	// Technically this is no longer needed as Bitcoin Cash has a different
	// network magic than Bitcoin Core so connections should not be possible.
	SFNodeBitcoinCash

	// SFNodeGraphene is a flag used to indicate a peer supports graphene block relay.
	SFNodeGraphene

	// SFNodeWeakBlocks is a flag used to indicate a peer supports the weak block protocol.
	SFNodeWeakBlocks

	// SFNodeCF is a flag used to indicate a peer supports committed
	// filters (CFs).
	SFNodeCF

	// SFNodeXThinner is a placeholder for the xthinner block compression protocol being
	// developed by Johnathan Toomim.
	SFNodeXThinner

	// SFNodeNetworkLimited is used to indicate the node is a pruned node and may only
	// be capable of limited services. In particular it is only guaranteed to be able
	// to serve the last 288 blocks though it will respond to requests for earlier blocks
	// if it has them.
	SFNodeNetworkLimited
)

// Map of service flags back to their constant names for pretty printing.
var sfStrings = map[ServiceFlag]string{
	SFNodeNetwork:        "SFNodeNetwork",
	SFNodeGetUTXO:        "SFNodeGetUTXO",
	SFNodeBloom:          "SFNodeBloom",
	SFNodeWitness:        "SFNodeWitness",
	SFNodeXthin:          "SFNodeXthin",
	SFNodeBitcoinCash:    "SFNodeBitcoinCash",
	SFNodeGraphene:       "SFNodeGraphene",
	SFNodeWeakBlocks:     "SFNodeWeakBlocks",
	SFNodeCF:             "SFNodeCF",
	SFNodeXThinner:       "SFNodeXThinner",
	SFNodeNetworkLimited: "SFNodeNetworkLimited",
}

// orderedSFStrings is an ordered list of service flags from highest to
// lowest.
var orderedSFStrings = []ServiceFlag{
	SFNodeNetwork,
	SFNodeGetUTXO,
	SFNodeBloom,
	SFNodeWitness,
	SFNodeXthin,
	SFNodeBitcoinCash,
	SFNodeGraphene,
	SFNodeWeakBlocks,
	SFNodeCF,
	SFNodeXThinner,
	SFNodeNetworkLimited,
}

// String returns the ServiceFlag in human-readable form.
func (f ServiceFlag) String() string {
	// No flags are set.
	if f == 0 {
		return "0x0"
	}

	// Add individual bit flags.
	s := ""
	for _, flag := range orderedSFStrings {
		if f&flag == flag {
			s += sfStrings[flag] + "|"
			f -= flag
		}
	}

	// Add any remaining flags which aren't accounted for as hex.
	s = strings.TrimRight(s, "|")
	if f != 0 {
		s += "|0x" + strconv.FormatUint(uint64(f), 16)
	}
	s = strings.TrimLeft(s, "|")
	return s
}

// BitcoinNet represents which bitcoin network a message belongs to.
type BitcoinNet uint32

// Constants used to indicate the message bitcoin network.  They can also be
// used to seek to the next message when a stream's state is unknown, but
// this package does not provide that functionality since it's generally a
// better idea to simply disconnect clients that are misbehaving over TCP.
const (
	// MainNet represents the main bitcoin network.
	MainNet BitcoinNet = 0xe8f3e1e3

	// TestNet represents the regression test network.
	RegTestNet BitcoinNet = 0xfabfb5da

	// TestNet represents the test network (version 3).
	TestNet BitcoinNet = 0xf4f3e5f4

	// SimNet represents the simulation test network.
	SimNet BitcoinNet = 0x12141c16
)

// bnStrings is a map of bitcoin networks back to their constant names for
// pretty printing.
var bnStrings = map[BitcoinNet]string{
	MainNet:    "MainNet",
	RegTestNet: "RegTestNet",
	TestNet:    "TestNet",
	SimNet:     "SimNet",
}

// String returns the BitcoinNet in human-readable form.
func (n BitcoinNet) String() string {
	if s, ok := bnStrings[n]; ok {
		return s
	}

	return fmt.Sprintf("Unknown BitcoinNet (%d)", uint32(n))
}