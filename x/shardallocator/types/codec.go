package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterNode{}, "reflexnet/RegisterNode", nil)
	cdc.RegisterConcrete(&MsgUnregisterNode{}, "reflexnet/UnregisterNode", nil)
	cdc.RegisterConcrete(&MsgAssignShard{}, "reflexnet/AssignShard", nil)
	cdc.RegisterConcrete(&MsgReplaceShard{}, "reflexnet/ReplaceShard", nil)
	cdc.RegisterConcrete(&MsgUpdateNodeStatus{}, "reflexnet/UpdateNodeStatus", nil)
}

// RegisterInterfaces registers the interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterNode{},
		&MsgUnregisterNode{},
		&MsgAssignShard{},
		&MsgReplaceShard{},
		&MsgUpdateNodeStatus{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(Amino)
	Amino.Seal()
}

