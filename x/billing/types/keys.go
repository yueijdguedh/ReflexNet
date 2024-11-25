package types

const (
	ModuleName   = "billing"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	MemStoreKey  = "mem_billing"
	QuerierRoute = ModuleName
)

var (
	BillingRecordKey       = []byte{0x01}
	PaymentDistributionKey = []byte{0x02}
	ParamsKey              = []byte{0x03}
)

func GetBillingRecordKey(requestID string) []byte {
	return append(BillingRecordKey, []byte(requestID)...)
}

func GetPaymentDistributionKey(requestID string) []byte {
	return append(PaymentDistributionKey, []byte(requestID)...)
}

