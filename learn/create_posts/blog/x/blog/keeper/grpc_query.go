package keeper

import (
	"github.com/papillon6814/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
