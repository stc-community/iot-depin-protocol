package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stc-community/iot-depin-protocol/testutil/keeper"
	"github.com/stc-community/iot-depin-protocol/testutil/nullify"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/keeper"
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDevice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Device {
	items := make([]types.Device, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetDevice(ctx, items[i])
	}
	return items
}

func TestDeviceGet(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDevice(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceRemove(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDevice(ctx,
			item.Address,
		)
		_, found := keeper.GetDevice(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestDeviceGetAll(t *testing.T) {
	keeper, ctx := keepertest.IotdepinprotocolKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDevice(ctx)),
	)
}
