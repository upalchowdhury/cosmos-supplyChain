package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/upalchowdhury/supplychain/testutil/keeper"
	"github.com/upalchowdhury/supplychain/testutil/nullify"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActorscontractQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActorscontract(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActorscontractRequest
		response *types.QueryGetActorscontractResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActorscontractRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetActorscontractResponse{Actorscontract: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActorscontractRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetActorscontractResponse{Actorscontract: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActorscontractRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Actorscontract(wctx, tc.request)
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

func TestActorscontractQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActorscontract(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActorscontractRequest {
		return &types.QueryAllActorscontractRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActorscontractAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Actorscontract), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Actorscontract),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActorscontractAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Actorscontract), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Actorscontract),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActorscontractAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Actorscontract),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActorscontractAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
