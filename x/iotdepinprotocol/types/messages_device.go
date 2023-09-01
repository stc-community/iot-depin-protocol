package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDevice = "create_device"
	TypeMsgUpdateDevice = "update_device"
	TypeMsgDeleteDevice = "delete_device"
)

var _ sdk.Msg = &MsgCreateDevice{}

func NewMsgCreateDevice(
	creator string,
	deviceName string,
	address string,
	value string,

) *MsgCreateDevice {
	return &MsgCreateDevice{
		Creator:    creator,
		DeviceName: deviceName,
		Address:    address,
		Value:      value,
	}
}

func (msg *MsgCreateDevice) Route() string {
	return RouterKey
}

func (msg *MsgCreateDevice) Type() string {
	return TypeMsgCreateDevice
}

func (msg *MsgCreateDevice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDevice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDevice{}

func NewMsgUpdateDevice(
	creator string,
	deviceName string,
	address string,
	value string,

) *MsgUpdateDevice {
	return &MsgUpdateDevice{
		Creator:    creator,
		DeviceName: deviceName,
		Address:    address,
		Value:      value,
	}
}

func (msg *MsgUpdateDevice) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDevice) Type() string {
	return TypeMsgUpdateDevice
}

func (msg *MsgUpdateDevice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDevice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDevice{}

func NewMsgDeleteDevice(
	creator string,
	deviceName string,

) *MsgDeleteDevice {
	return &MsgDeleteDevice{
		Creator:    creator,
		DeviceName: deviceName,
	}
}
func (msg *MsgDeleteDevice) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDevice) Type() string {
	return TypeMsgDeleteDevice
}

func (msg *MsgDeleteDevice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDevice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
