package types

const (
	ModuleName   = "inferencegateway"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	MemStoreKey  = "mem_inferencegateway"
	QuerierRoute = ModuleName
)

var (
	InferenceRequestKey  = []byte{0x01}
	InferenceResponseKey = []byte{0x02}
	ParamsKey            = []byte{0x03}
	NonceKey             = []byte{0x04}
)

func GetInferenceRequestKey(requestID string) []byte {
	return append(InferenceRequestKey, []byte(requestID)...)
}

func GetInferenceResponseKey(requestID string) []byte {
	return append(InferenceResponseKey, []byte(requestID)...)
}

func GetNonceKey(address string) []byte {
	return append(NonceKey, []byte(address)...)
}

