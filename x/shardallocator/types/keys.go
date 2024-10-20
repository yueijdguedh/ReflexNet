package types

const (
	// ModuleName defines the module name
	ModuleName = "shardallocator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_shardallocator"

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// Store key prefixes
var (
	// ShardAssignmentKey is the prefix for storing shard assignments
	ShardAssignmentKey = []byte{0x01}

	// NodeInfoKey is the prefix for storing node information
	NodeInfoKey = []byte{0x02}

	// NodeByStatusKey is the prefix for indexing nodes by status
	NodeByStatusKey = []byte{0x03}

	// ParamsKey is the key for storing module parameters
	ParamsKey = []byte{0x04}
)

// GetShardAssignmentKey returns the store key for a shard assignment
func GetShardAssignmentKey(modelID uint64, shardIndex uint32) []byte {
	return append(append(ShardAssignmentKey, Uint64ToBytes(modelID)...), Uint32ToBytes(shardIndex)...)
}

// GetNodeInfoKey returns the store key for node information
func GetNodeInfoKey(nodeAddress string) []byte {
	return append(NodeInfoKey, []byte(nodeAddress)...)
}

// Uint64ToBytes converts uint64 to bytes
func Uint64ToBytes(val uint64) []byte {
	b := make([]byte, 8)
	for i := 0; i < 8; i++ {
		b[i] = byte(val >> (8 * i))
	}
	return b
}

// Uint32ToBytes converts uint32 to bytes
func Uint32ToBytes(val uint32) []byte {
	b := make([]byte, 4)
	for i := 0; i < 4; i++ {
		b[i] = byte(val >> (8 * i))
	}
	return b
}

