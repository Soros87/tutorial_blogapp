package keeper

import (
	"blogapp/x/blogapp/types"
)

var _ types.QueryServer = Keeper{}
