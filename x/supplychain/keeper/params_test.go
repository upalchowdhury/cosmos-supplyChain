package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SupplychainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
