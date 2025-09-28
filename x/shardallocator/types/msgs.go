package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *MsgRegisterNode) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid node address: %s", err)
	}
	
	if !m.StakeAmount.IsValid() || m.StakeAmount.IsZero() {
		return sdkerrors.ErrInvalidCoins.Wrap("stake amount must be positive")
	}
	
	return nil
}

func (m *MsgRegisterNode) GetSigners() []sdk.AccAddress {
	nodeAddr, err := sdk.AccAddressFromBech32(m.NodeAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{nodeAddr}
}

func (m *MsgUnregisterNode) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid node address: %s", err)
	}
	return nil
}

func (m *MsgUnregisterNode) GetSigners() []sdk.AccAddress {
	nodeAddr, err := sdk.AccAddressFromBech32(m.NodeAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{nodeAddr}
}
