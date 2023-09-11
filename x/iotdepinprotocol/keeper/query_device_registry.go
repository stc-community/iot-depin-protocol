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

func (k Keeper) DeviceRegistryAll(goCtx context.Context, req *types.QueryAllDeviceRegistryRequest) (*types.QueryAllDeviceRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var deviceRegistrys []types.DeviceRegistry
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	deviceRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.DeviceRegistryKeyPrefix))

	pageRes, err := query.Paginate(deviceRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var deviceRegistry types.DeviceRegistry
		if err := k.cdc.Unmarshal(value, &deviceRegistry); err != nil {
			return err
		}

		deviceRegistrys = append(deviceRegistrys, deviceRegistry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDeviceRegistryResponse{DeviceRegistry: deviceRegistrys, Pagination: pageRes}, nil
}

func (k Keeper) DeviceRegistry(goCtx context.Context, req *types.QueryGetDeviceRegistryRequest) (*types.QueryGetDeviceRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDeviceRegistry(
		ctx,
		req.Mid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeviceRegistryResponse{DeviceRegistry: val}, nil
}
