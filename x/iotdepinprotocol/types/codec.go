package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateKv{}, "iotdepinprotocol/CreateKv", nil)
	cdc.RegisterConcrete(&MsgUpdateKv{}, "iotdepinprotocol/UpdateKv", nil)
	cdc.RegisterConcrete(&MsgDeleteKv{}, "iotdepinprotocol/DeleteKv", nil)
	cdc.RegisterConcrete(&MsgCreateDevice{}, "iotdepinprotocol/CreateDevice", nil)
	cdc.RegisterConcrete(&MsgUpdateDevice{}, "iotdepinprotocol/UpdateDevice", nil)
	cdc.RegisterConcrete(&MsgDeleteDevice{}, "iotdepinprotocol/DeleteDevice", nil)
	cdc.RegisterConcrete(&MsgOracleOperator{}, "iotdepinprotocol/OracleOperator", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKv{},
		&MsgUpdateKv{},
		&MsgDeleteKv{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDevice{},
		&MsgUpdateDevice{},
		&MsgDeleteDevice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOracleOperator{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
