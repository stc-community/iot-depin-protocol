package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDeviceRegistry = "create_device_registry"
	TypeMsgUpdateDeviceRegistry = "update_device_registry"
	TypeMsgDeleteDeviceRegistry = "delete_device_registry"
)

var _ sdk.Msg = &MsgCreateDeviceRegistry{}

func NewMsgCreateDeviceRegistry(
	creator string,
	mid string,
	metaData string,

) *MsgCreateDeviceRegistry {
	return &MsgCreateDeviceRegistry{
		Creator:  creator,
		Mid:      mid,
		MetaData: metaData,
	}
}

func (msg *MsgCreateDeviceRegistry) Route() string {
	return RouterKey
}

func (msg *MsgCreateDeviceRegistry) Type() string {
	return TypeMsgCreateDeviceRegistry
}

func (msg *MsgCreateDeviceRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDeviceRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDeviceRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDeviceRegistry{}

func NewMsgUpdateDeviceRegistry(
	creator string,
	mid string,
	metaData string,

) *MsgUpdateDeviceRegistry {
	return &MsgUpdateDeviceRegistry{
		Creator:  creator,
		Mid:      mid,
		MetaData: metaData,
	}
}

func (msg *MsgUpdateDeviceRegistry) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDeviceRegistry) Type() string {
	return TypeMsgUpdateDeviceRegistry
}

func (msg *MsgUpdateDeviceRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDeviceRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDeviceRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDeviceRegistry{}

func NewMsgDeleteDeviceRegistry(
	creator string,
	mid string,

) *MsgDeleteDeviceRegistry {
	return &MsgDeleteDeviceRegistry{
		Creator: creator,
		Mid:     mid,
	}
}
func (msg *MsgDeleteDeviceRegistry) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDeviceRegistry) Type() string {
	return TypeMsgDeleteDeviceRegistry
}

func (msg *MsgDeleteDeviceRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDeviceRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDeviceRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
