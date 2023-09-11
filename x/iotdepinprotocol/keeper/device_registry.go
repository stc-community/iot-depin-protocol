package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// SetDeviceRegistry set a specific deviceRegistry in the store from its index
func (k Keeper) SetDeviceRegistry(ctx sdk.Context, deviceRegistry types.DeviceRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&deviceRegistry)
	store.Set(types.DeviceRegistryKey(
		deviceRegistry.Mid,
	), b)
}

// GetDeviceRegistry returns a deviceRegistry from its index
func (k Keeper) GetDeviceRegistry(
	ctx sdk.Context,
	mid string,

) (val types.DeviceRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceRegistryKeyPrefix))

	b := store.Get(types.DeviceRegistryKey(
		mid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeviceRegistry removes a deviceRegistry from the store
func (k Keeper) RemoveDeviceRegistry(
	ctx sdk.Context,
	mid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceRegistryKeyPrefix))
	store.Delete(types.DeviceRegistryKey(
		mid,
	))
}

// GetAllDeviceRegistry returns all deviceRegistry
func (k Keeper) GetAllDeviceRegistry(ctx sdk.Context) (list []types.DeviceRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DeviceRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
