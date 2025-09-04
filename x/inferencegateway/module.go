package inferencegateway

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/yueijdguedh/ReflexNet/x/inferencegateway/keeper"
)

type AppModule struct {
	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{keeper: keeper}
}
