package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func (k msgServer) CreateTran(goCtx context.Context, msg *types.MsgCreateTran) (*types.MsgCreateTranResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTranResponse{}, nil
}
