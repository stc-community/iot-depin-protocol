package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOracleOperator = "oracle_operator"

var _ sdk.Msg = &MsgOracleOperator{}

func NewMsgOracleOperator(creator string, reqId string, operatorType string, data string) *MsgOracleOperator {
	return &MsgOracleOperator{
		Creator:      creator,
		ReqId:        reqId,
		OperatorType: operatorType,
		Data:         data,
	}
}

func (msg *MsgOracleOperator) Route() string {
	return RouterKey
}

func (msg *MsgOracleOperator) Type() string {
	return TypeMsgOracleOperator
}

func (msg *MsgOracleOperator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOracleOperator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOracleOperator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
