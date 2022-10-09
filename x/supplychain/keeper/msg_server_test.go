package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SupplychainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
