package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *MsgSubmitInferenceRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Requester); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid requester address: %s", err)
	}
	
	if m.ModelId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("model ID cannot be zero")
	}
	
	if m.InputData == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("input data cannot be empty")
	}
	
	return nil
}

func (m *MsgSubmitInferenceRequest) GetSigners() []sdk.AccAddress {
	requester, err := sdk.AccAddressFromBech32(m.Requester)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{requester}
}
