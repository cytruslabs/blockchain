package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAd{}, "zestchain/CreateAd", nil)
	cdc.RegisterConcrete(&MsgPayView{}, "zestchain/PayView", nil)
	cdc.RegisterConcrete(&MsgPayClick{}, "zestchain/PayClick", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAd{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPayView{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPayClick{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
