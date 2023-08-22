package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DeviceAll(goCtx context.Context, req *types.QueryAllDeviceRequest) (*types.QueryAllDeviceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var devices []types.Device
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	deviceStore := prefix.NewStore(store, types.KeyPrefix(types.DeviceKeyPrefix))

	pageRes, err := query.Paginate(deviceStore, req.Pagination, func(key []byte, value []byte) error {
		var device types.Device
		if err := k.cdc.Unmarshal(value, &device); err != nil {
			return err
		}

		devices = append(devices, device)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDeviceResponse{Device: devices, Pagination: pageRes}, nil
}

func (k Keeper) Device(goCtx context.Context, req *types.QueryGetDeviceRequest) (*types.QueryGetDeviceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDevice(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeviceResponse{Device: val}, nil
}
