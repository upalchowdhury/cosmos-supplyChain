package supplychain

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/upalchowdhury/supplychain/x/supplychain/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateTran:
			res, err := msgServer.CreateTran(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateContract:
			res, err := msgServer.CreateContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			case *types.MsgCommitContract:
					res, err := msgServer.CommitContract(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgApproveContract:
					res, err := msgServer.ApproveContract(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgShipOrder:
					res, err := msgServer.ShipOrder(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgOrderDelivered:
					res, err := msgServer.OrderDelivered(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
case *types.MsgCancelOrder:
					res, err := msgServer.CancelOrder(sdk.WrapSDKContext(ctx), msg)
					return sdk.WrapServiceResult(ctx, res, err)
// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
