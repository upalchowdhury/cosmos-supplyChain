package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewcontractAll(c context.Context, req *types.QueryAllNewcontractRequest) (*types.QueryAllNewcontractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var newcontracts []types.Newcontract
	ctx := sdk.UnwrapSDKContext(c)

	prefixStoreKey := types.NewcontractKey(req.DealId)

	store := ctx.KVStore(k.storeKey)
	newcontractStore := prefix.NewStore(store, prefixStoreKey)

	pageRes, err := query.Paginate(newcontractStore, req.Pagination, func(key []byte, value []byte) error {
		var newcontract types.Newcontract
		if err := k.cdc.Unmarshal(value, &newcontract); err != nil {
			return err
		}

		newcontracts = append(newcontracts, newcontract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNewcontractResponse{Newcontract: newcontracts, Pagination: pageRes}, nil
}

func (k Keeper) Newcontract(c context.Context, req *types.QueryGetNewcontractRequest) (*types.QueryGetNewcontractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNewcontract(
		ctx,
		req.DealId,
		req.ContractId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNewcontractResponse{Newcontract: val}, nil
	// next register query server on module
}
