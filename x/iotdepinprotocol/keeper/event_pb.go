package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

// GetEventPbCount get the total number of eventPb
func (k Keeper) GetEventPbCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EventPbCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEventPbCount set the total number of eventPb
func (k Keeper) SetEventPbCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EventPbCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEventPb appends a eventPb in the store with a new id and update the count
func (k Keeper) AppendEventPb(
	ctx sdk.Context,
	eventPb types.EventPb,
) uint64 {
	// Create the eventPb
	count := k.GetEventPbCount(ctx)

	// Set the ID of the appended value
	eventPb.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKey))
	appendedValue := k.cdc.MustMarshal(&eventPb)
	store.Set(GetEventPbIDBytes(eventPb.Id), appendedValue)

	// Update eventPb count
	k.SetEventPbCount(ctx, count+1)

	return count
}

// SetEventPb set a specific eventPb in the store
func (k Keeper) SetEventPb(ctx sdk.Context, eventPb types.EventPb) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKey))
	b := k.cdc.MustMarshal(&eventPb)
	store.Set(GetEventPbIDBytes(eventPb.Id), b)
}

// GetEventPb returns a eventPb from its id
func (k Keeper) GetEventPb(ctx sdk.Context, id uint64) (val types.EventPb, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKey))
	b := store.Get(GetEventPbIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEventPb removes a eventPb from the store
func (k Keeper) RemoveEventPb(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKey))
	store.Delete(GetEventPbIDBytes(id))
}

// GetAllEventPb returns all eventPb
func (k Keeper) GetAllEventPb(ctx sdk.Context) (list []types.EventPb) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EventPbKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EventPb
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEventPbIDBytes returns the byte representation of the ID
func GetEventPbIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEventPbIDFromBytes returns ID in uint64 format from a byte array
func GetEventPbIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
