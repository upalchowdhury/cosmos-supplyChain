package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func (k msgServer) ShipOrder(goCtx context.Context, msg *types.MsgShipOrder) (*types.MsgShipOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	deal, found := k.Keeper.GetNewDeal(ctx, msg.DealId)
	if !found {
		return nil, types.ErrDealNotFound
	}
	// validate
	if msg.Creator != deal.Vendor {
		return nil, types.ErrInvalidVendor
	}
	contract, found := k.Keeper.GetNewcontract(ctx, msg.DealId,
		msg.ContractId)
	if !found {
		return nil, types.ErrContractNotFound
	}
	if contract.Status != types.APPROVED || contract.Status == types.COMPLETED {
		return nil, types.ErrNotApprovedOrCompleted
	}
	startTime, err := time.Parse(types.TIME_FORMAT, contract.StartTime)
	if err != nil {
		panic("invalid start time")
	}
	// calculate shipping delay
	shippingExpectedTime := startTime.Add(time.Duration(contract.VendorETA))
	shippingActualTime := ctx.BlockTime()
	if shippingActualTime.After(shippingExpectedTime) {
		shippingTimeDelay := shippingActualTime.Sub(shippingExpectedTime).Minutes()
		contract.ShippingDelay = uint32(shippingTimeDelay)
	}

	contract.Status = types.INDELIVERY
	k.Keeper.SetNewcontract(ctx, contract)
	ctx.GasMeter().ConsumeGas(types.PROCESS_GAS, "Order shipped")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.INDELIVERY),
			sdk.NewAttribute(types.IDVALUE, contract.ContractId),
		),
	)
	return &types.MsgShipOrderResponse{IdValue: contract.ContractId, ContractStatus: contract.Status}, nil
}
