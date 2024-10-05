package types

const (
	// ModuleName defines the module name
	ModuleName = "modelregistry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_modelregistry"

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// Store key prefixes
var (
	// ModelKey is the prefix for storing models by ID
	ModelKey = []byte{0x01}

	// ModelByOwnerKey is the prefix for indexing models by owner
	ModelByOwnerKey = []byte{0x02}

	// NextModelIDKey is the key for storing the next model ID
	NextModelIDKey = []byte{0x03}

	// ParamsKey is the key for storing module parameters
	ParamsKey = []byte{0x04}
)

// GetModelKey returns the store key for a model by ID
func GetModelKey(modelID uint64) []byte {
	return append(ModelKey, Uint64ToBytes(modelID)...)
}

// GetModelByOwnerKey returns the store key for indexing models by owner
func GetModelByOwnerKey(owner string, modelID uint64) []byte {
	return append(append(ModelByOwnerKey, []byte(owner)...), Uint64ToBytes(modelID)...)
}

// Uint64ToBytes converts uint64 to bytes
func Uint64ToBytes(val uint64) []byte {
	b := make([]byte, 8)
	for i := 0; i < 8; i++ {
		b[i] = byte(val >> (8 * i))
	}
	return b
}

// BytesToUint64 converts bytes to uint64
func BytesToUint64(b []byte) uint64 {
	var val uint64
	for i := 0; i < 8 && i < len(b); i++ {
		val |= uint64(b[i]) << (8 * i)
	}
	return val
}

