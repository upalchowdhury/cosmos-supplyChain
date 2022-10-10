package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// SetActorscontract set a specific actorscontract in the store from its index
func (k Keeper) SetActorscontract(ctx sdk.Context, actorscontract types.Actorscontract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActorscontractKeyPrefix))
	b := k.cdc.MustMarshal(&actorscontract)
	store.Set(types.ActorscontractKey(
		actorscontract.Index,
	), b)
}

// GetActorscontract returns a actorscontract from its index
func (k Keeper) GetActorscontract(
	ctx sdk.Context,
	index string,

) (val types.Actorscontract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActorscontractKeyPrefix))

	b := store.Get(types.ActorscontractKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActorscontract removes a actorscontract from the store
func (k Keeper) RemoveActorscontract(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActorscontractKeyPrefix))
	store.Delete(types.ActorscontractKey(
		index,
	))
}

// GetAllActorscontract returns all actorscontract
func (k Keeper) GetAllActorscontract(ctx sdk.Context) (list []types.Actorscontract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActorscontractKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Actorscontract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
