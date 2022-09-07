package keeper

import (
	"github.com/Zireael26/hello-blockchain-world/x/helloblockchainworld/types"
)

var _ types.QueryServer = Keeper{}
