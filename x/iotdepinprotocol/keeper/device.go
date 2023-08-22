package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// SetDevice set a specific device in the store from its index
func (k Keeper) SetDevice(ctx sdk.Context, device types.Device) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix+device.Creator))
	b := k.cdc.MustMarshal(&device)
	store.Set(types.DeviceKey(
		device.Address,
	), b)
}

// GetDevice returns a device from its index
func (k Keeper) GetDevice(
	ctx sdk.Context,
	address string,
	creator string,
) (val types.Device, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix+creator))

	b := store.Get(types.DeviceKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDevice removes a device from the store
func (k Keeper) RemoveDevice(
	ctx sdk.Context,
	address string,
	creator string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix+creator))
	store.Delete(types.DeviceKey(
		address,
	))
}

// GetAllDevice returns all device
func (k Keeper) GetAllDevice(ctx sdk.Context, creator string) (list []types.Device) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceKeyPrefix+creator))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Device
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
