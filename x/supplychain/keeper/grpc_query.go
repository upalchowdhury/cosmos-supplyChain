package keeper

import (
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

var _ types.QueryServer = Keeper{}
