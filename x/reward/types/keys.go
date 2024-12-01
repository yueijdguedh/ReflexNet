package types

const (
	ModuleName   = "reward"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	MemStoreKey  = "mem_reward"
	QuerierRoute = ModuleName
)

var (
	RewardPoolKey     = []byte{0x01}
	NodeRewardKey     = []byte{0x02}
	SlashingEventKey  = []byte{0x03}
	ParamsKey         = []byte{0x04}
)

func GetNodeRewardKey(nodeAddress string) []byte {
	return append(NodeRewardKey, []byte(nodeAddress)...)
}

func GetSlashingEventKey(nodeAddress string, height int64) []byte {
	key := append(SlashingEventKey, []byte(nodeAddress)...)
	return append(key, Int64ToBytes(height)...)
}

func Int64ToBytes(val int64) []byte {
	b := make([]byte, 8)
	for i := 0; i < 8; i++ {
		b[i] = byte(val >> (8 * i))
	}
	return b
}

