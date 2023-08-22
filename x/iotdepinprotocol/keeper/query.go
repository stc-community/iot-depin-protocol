package keeper

import (
	"github.com/stc-community/iot-depin-protocol/x/iotdepinprotocol/types"
)

var _ types.QueryServer = Keeper{}
