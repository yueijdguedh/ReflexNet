package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgRegisterModel{}
	_ sdk.Msg = &MsgUpdateModelVersion{}
	_ sdk.Msg = &MsgUpdateModelStatus{}
)

// ValidateBasic performs basic validation on MsgRegisterModel
func (m *MsgRegisterModel) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}

	if m.Name == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("model name cannot be empty")
	}

	if len(m.Name) > 256 {
		return ErrModelNameTooLong
	}

	if m.MetadataCid == "" {
		return ErrInvalidMetadataCID
	}

	if m.ShardCount == 0 {
		return ErrInvalidShardCount
	}

	if m.Version == "" {
		return ErrInvalidVersion
	}

	return nil
}

// GetSigners returns the signers of MsgRegisterModel
func (m *MsgRegisterModel) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// ValidateBasic performs basic validation on MsgUpdateModelVersion
func (m *MsgUpdateModelVersion) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}

	if m.ModelId == 0 {
		return ErrInvalidModelID
	}

	if m.NewVersion == "" {
		return ErrInvalidVersion
	}

	if m.NewMetadataCid == "" {
		return ErrInvalidMetadataCID
	}

	return nil
}

// GetSigners returns the signers of MsgUpdateModelVersion
func (m *MsgUpdateModelVersion) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// ValidateBasic performs basic validation on MsgUpdateModelStatus
func (m *MsgUpdateModelStatus) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}

	if m.ModelId == 0 {
		return ErrInvalidModelID
	}

	if m.NewStatus == ModelStatus_MODEL_STATUS_UNSPECIFIED {
		return ErrInvalidModelStatus
	}

	return nil
}

// GetSigners returns the signers of MsgUpdateModelStatus
func (m *MsgUpdateModelStatus) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(m.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

