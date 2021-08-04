package keeper

import (
	"github.com/papillon6814/lotus-ibc/x/lotusibc/types"
)

var _ types.QueryServer = Keeper{}
