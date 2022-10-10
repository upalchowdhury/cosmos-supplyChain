package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// SetConCounter set conCounter in the store
func (k Keeper) SetConCounter(ctx sdk.Context, conCounter types.ConCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConCounterKey))
	b := k.cdc.MustMarshal(&conCounter)
	store.Set([]byte{0}, b)
}

// GetConCounter returns conCounter
func (k Keeper) GetConCounter(ctx sdk.Context) (val types.ConCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConCounterKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveConCounter removes conCounter from the store
func (k Keeper) RemoveConCounter(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConCounterKey))
	store.Delete([]byte{0})
}
