package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) msgServer {
	return msgServer{Keeper: keeper}
}
