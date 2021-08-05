package keeper

import (
	"github.com/papillon6814/lotus-ibc/x/blog/types"
)

var _ types.QueryServer = Keeper{}
