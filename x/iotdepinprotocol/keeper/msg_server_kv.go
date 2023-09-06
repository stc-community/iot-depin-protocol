package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

func (k msgServer) CreateKv(goCtx context.Context, msg *types.MsgCreateKv) (*types.MsgCreateKvResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 验证设备
	deviceFound, isFound := k.GetDevice(ctx, msg.DeviceName)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "device not fount")
	}
	if deviceFound.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only device creator can create kv")
	}

	// Check if the value already exists
	_, isFound = k.GetKv(
		ctx,
		msg.Index,
		msg.DeviceName,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var kv = types.Kv{
		Creator:    msg.Creator,
		Index:      msg.Index,
		Value:      msg.Value,
		DeviceName: msg.DeviceName,
	}

	k.SetKv(
		ctx,
		kv,
	)
	return &types.MsgCreateKvResponse{}, nil
}

func (k msgServer) UpdateKv(goCtx context.Context, msg *types.MsgUpdateKv) (*types.MsgUpdateKvResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKv(
		ctx,
		msg.Index,
		msg.DeviceName,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var kv = types.Kv{
		Creator:    msg.Creator,
		Index:      msg.Index,
		Value:      msg.Value,
		DeviceName: msg.DeviceName,
	}

	k.SetKv(ctx, kv)

	return &types.MsgUpdateKvResponse{}, nil
}

func (k msgServer) DeleteKv(goCtx context.Context, msg *types.MsgDeleteKv) (*types.MsgDeleteKvResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKv(
		ctx,
		msg.Index,
		msg.DeviceName,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKv(
		ctx,
		msg.Index,
		msg.DeviceName,
	)

	return &types.MsgDeleteKvResponse{}, nil
}
