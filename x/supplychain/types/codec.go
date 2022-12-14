package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateTran{}, "supplychain/CreateTran", nil)
	cdc.RegisterConcrete(&MsgCreateContract{}, "supplychain/CreateContract", nil)
	cdc.RegisterConcrete(&MsgCommitContract{}, "supplychain/CommitContract", nil)
cdc.RegisterConcrete(&MsgApproveContract{}, "supplychain/ApproveContract", nil)
cdc.RegisterConcrete(&MsgShipOrder{}, "supplychain/ShipOrder", nil)
cdc.RegisterConcrete(&MsgOrderDelivered{}, "supplychain/OrderDelivered", nil)
cdc.RegisterConcrete(&MsgCancelOrder{}, "supplychain/CancelOrder", nil)
// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTran{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCommitContract{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgApproveContract{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgShipOrder{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgOrderDelivered{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCancelOrder{},
)
// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
