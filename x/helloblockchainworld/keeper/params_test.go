package keeper_test

import (
	"testing"

	testkeeper "github.com/Zireael26/hello-blockchain-world/testutil/keeper"
	"github.com/Zireael26/hello-blockchain-world/x/helloblockchainworld/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HelloblockchainworldKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
