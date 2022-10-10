package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/testutil/nullify"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func TestConCounterQuery(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestConCounter(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetConCounterRequest
		response *types.QueryGetConCounterResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetConCounterRequest{},
			response: &types.QueryGetConCounterResponse{ConCounter: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ConCounter(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
