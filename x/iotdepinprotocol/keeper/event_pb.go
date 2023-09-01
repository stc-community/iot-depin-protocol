package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// SetEventPb set a specific eventPb in the store from its index
func (k Keeper) SetEventPb(ctx sdk.Context, eventPb types.EventPb) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKeyPrefix+eventPb.DeviceName))
	b := k.cdc.MustMarshal(&eventPb)
	store.Set(types.EventPbKey(
		eventPb.Index,
	), b)
}

// GetEventPb returns a eventPb from its index
func (k Keeper) GetEventPb(
	ctx sdk.Context,
	index string,
	deviceName string,

) (val types.EventPb, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKeyPrefix+deviceName))

	b := store.Get(types.EventPbKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEventPb removes a eventPb from the store
func (k Keeper) RemoveEventPb(
	ctx sdk.Context,
	index string,
	deviceName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKeyPrefix+deviceName))
	store.Delete(types.EventPbKey(
		index,
	))
}

// GetAllEventPb returns all eventPb
func (k Keeper) GetAllEventPb(ctx sdk.Context) (list []types.EventPb) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EventPb
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
