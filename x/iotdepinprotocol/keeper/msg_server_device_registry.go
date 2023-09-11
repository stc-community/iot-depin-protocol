package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) CreateDeviceRegistry(goCtx context.Context, msg *types.MsgCreateDeviceRegistry) (*types.MsgCreateDeviceRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDeviceRegistry(
		ctx,
		msg.Mid,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var deviceRegistry = types.DeviceRegistry{
		Creator:  msg.Creator,
		Mid:      msg.Mid,
		MetaData: msg.MetaData,
	}

	k.SetDeviceRegistry(
		ctx,
		deviceRegistry,
	)
	return &types.MsgCreateDeviceRegistryResponse{}, nil
}

func (k msgServer) UpdateDeviceRegistry(goCtx context.Context, msg *types.MsgUpdateDeviceRegistry) (*types.MsgUpdateDeviceRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeviceRegistry(
		ctx,
		msg.Mid,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var deviceRegistry = types.DeviceRegistry{
		Creator:  msg.Creator,
		Mid:      msg.Mid,
		MetaData: msg.MetaData,
	}

	k.SetDeviceRegistry(ctx, deviceRegistry)

	return &types.MsgUpdateDeviceRegistryResponse{}, nil
}

func (k msgServer) DeleteDeviceRegistry(goCtx context.Context, msg *types.MsgDeleteDeviceRegistry) (*types.MsgDeleteDeviceRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeviceRegistry(
		ctx,
		msg.Mid,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDeviceRegistry(
		ctx,
		msg.Mid,
	)

	return &types.MsgDeleteDeviceRegistryResponse{}, nil
}
