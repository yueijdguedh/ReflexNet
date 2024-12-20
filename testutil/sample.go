package testutil

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := secp256k1.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// ModelID returns a sample model ID
func ModelID() uint64 {
	return 1
}

// ShardCount returns a sample shard count
func ShardCount() uint32 {
	return 10
}

// MetadataCID returns a sample IPFS CID
func MetadataCID() string {
	return "QmXxxx1234567890abcdef"
}

// Version returns a sample version string
func Version() string {
	return "1.0.0"
}

