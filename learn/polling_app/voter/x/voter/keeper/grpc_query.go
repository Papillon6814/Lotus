package keeper

import (
	"github.com/papillon6814/voter/x/voter/types"
)

var _ types.QueryServer = Keeper{}
