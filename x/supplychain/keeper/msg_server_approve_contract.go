package keeper

import (
	"context"

    "github.com/upalchowdhury/supplychain/x/supplychain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) ApproveContract(goCtx context.Context,  msg *types.MsgApproveContract) (*types.MsgApproveContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
   // ApproveContract is the tx handler for handling approve contract messages
func (k msgServer) ApproveContract(goCtx context.Context, msg *types.MsgApproveContract) (*types.MsgApproveContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	contract, found := k.Keeper.GetNewContract(ctx, msg.DealId,  msg.ContractId)
	if !found {
	 return nil, types.ErrContractNotFound
	}
	func (msg *MsgApproveContract) DealHandlerValidation(goCtx context.Context, contract *NewContract) error {
		if msg.Creator != contract.Consumer {
		 return ErrInvalidConsumer
		}
		if contract.Status != COMMITTED {
		 return ErrNotCommitted
		}
		return nil
		}
	// handle validation 
	err := msg.DealHandlerValidation(goCtx, &contract)
	if err != nil {
	 return nil, err
	}
	expiryTime, err := time.Parse(types.TIME_FORMAT, contract.Expiry)
	if err != nil {
	 return nil, err
	}
	// don't process the expired contracts
	if ctx.BlockTime().After(expiryTime) {
	 return nil, types.ErrContractExpired
	}
	// store funds from user account to module escrow account and approve the contract
	consumerAddress, err := contract.GetConsumerAddress()
	if err != nil {
	panic("Invalid consumer address")
	}
	err = k.bank.SendCoinsFromAccountToModule(ctx, consumerAddress, types.ModuleName, sdk.NewCoins(contract.GetCoin(contract.Fees)))
	if err != nil {
	 return nil, sdkerrors.Wrapf(err, types.ErrPaymentFailed.Error())
	}
	contract.Status = types.APPROVED
	k.Keeper.SetNewContract(ctx, contract)
	// consume the gas to incentivize validators
	ctx.GasMeter().ConsumeGas(types.PROCESS_GAS, "Approve Contract")
	// emit custom event that clients can subscribe to
	ctx.EventManager().EmitEvent(
	sdk.NewEvent(sdk.EventTypeMessage,
	sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	sdk.NewAttribute(sdk.AttributeKeyAction, types.APPROVED),
	sdk.NewAttribute(types.IDVALUE, contract.ContractId),
	),
	)
	return &types.MsgApproveContractResponse{
	IdValue:        contract.ContractId,
	ContractStatus: contract.Status,
	}, nil
	}

	return &types.MsgApproveContractResponse{}, nil
}
