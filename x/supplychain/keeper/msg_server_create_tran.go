package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func (k msgServer) CreateTran(goCtx context.Context, msg *types.MsgCreateTran) (*types.MsgCreateTranResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	tranCounter, found := k.Keeper.GetTranCounter(ctx)

	if !found {
		panic("DealCounter not found")
	}

	dealId := strconv.FormatUint(tranCounter.IdValue, 10)

	newTran := types.Newcontract{
		DealId:     dealId,
		Owner:      msg.Creator,
		Vendor:     msg.Vendor,
		Commission: msg.Commission,
	}

	// validate before processing the message
	err := newTran.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetNewTran(ctx, newTran)
	tranCounter.IdValue++
	k.Keeper.SetTranCounter(ctx, tranCounter)
	// Set the new contract counter for a newly created deal
	contractCounter := types.ContractCounter{
		DealId:  dealId,
		IdValue: 0,
	}
	k.Keeper.SetContractCounter(ctx, contractCounter)
	ctx.GasMeter().ConsumeGas(types.CREATE_GAS, "Create Transaction")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.IDVALUE, dealId),
			sdk.NewAttribute(types.OWNER, newTran.Owner),
			sdk.NewAttribute(types.VENDOR, newTran.Vendor),
		),
	)

	return &types.MsgCreateTranResponse{
		IdValue: dealId,
	}, nil
}
