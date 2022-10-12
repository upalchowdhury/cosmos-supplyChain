package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// SetNewcontract set a specific newcontract in the store from its index
func (k Keeper) SetNewcontract(ctx sdk.Context, newcontract types.Newcontract) {

	key := types.NewcontractKey(newcontract.DealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), key)
	b := k.cdc.MustMarshal(&newcontract)

	store.Set(types.NewcontractKey(
		newcontract.Index,
	), b)
}

// GetNewcontract returns a newcontract from its index
func (k Keeper) GetNewcontract(
	ctx sdk.Context,
	index string,

) (val types.Newcontract, found bool) {

	storeKey := types.NewcontractKey(val.DealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)

	b := store.Get(types.NewcontractKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNewcontract removes a newcontract from the store
func (k Keeper) RemoveNewcontract(
	ctx sdk.Context,
	contractId string,
	dealId string,

) {
	storeKey := types.NewcontractKey(dealId)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	store.Delete(types.NewcontractKey(
		contractId,
	))
}

// GetAllNewcontract returns all newcontract
func (k Keeper) GetAllNewcontract(ctx sdk.Context, dealId string) (list []types.Newcontract) {

	storeKey := types.NewcontractKey(dealId)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Newcontract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
