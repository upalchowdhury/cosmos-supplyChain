package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

// SetContractCounter set contractCounter in the store
func (k Keeper) SetContractCounter(ctx sdk.Context, contractCounter types.ContractCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))
	b := k.cdc.MustMarshal(&contractCounter)
	store.Set([]byte{0}, b)
}

// GetContractCounter returns contractCounter
func (k Keeper) GetContractCounter(ctx sdk.Context) (val types.ContractCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContractCounter removes contractCounter from the store
func (k Keeper) RemoveContractCounter(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractCounterKey))
	store.Delete([]byte{0})
}
