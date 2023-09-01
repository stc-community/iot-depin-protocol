package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) CreateDevice(goCtx context.Context, msg *types.MsgCreateDevice) (*types.MsgCreateDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDevice(
		ctx,
		msg.DeviceName,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var device = types.Device{
		Creator:    msg.Creator,
		DeviceName: msg.DeviceName,
		Address:    msg.Address,
		Value:      msg.Value,
	}

	k.SetDevice(
		ctx,
		device,
	)
	return &types.MsgCreateDeviceResponse{}, nil
}

func (k msgServer) UpdateDevice(goCtx context.Context, msg *types.MsgUpdateDevice) (*types.MsgUpdateDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDevice(
		ctx,
		msg.DeviceName,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var device = types.Device{
		Creator:    msg.Creator,
		DeviceName: msg.DeviceName,
		Address:    msg.Address,
		Value:      msg.Value,
	}

	k.SetDevice(ctx, device)

	return &types.MsgUpdateDeviceResponse{}, nil
}

func (k msgServer) DeleteDevice(goCtx context.Context, msg *types.MsgDeleteDevice) (*types.MsgDeleteDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDevice(
		ctx,
		msg.DeviceName,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDevice(
		ctx,
		msg.DeviceName,
	)

	return &types.MsgDeleteDeviceResponse{}, nil
}
