package helloblockchainworld_test

import (
	"testing"

	keepertest "github.com/Zireael26/hello-blockchain-world/testutil/keeper"
	"github.com/Zireael26/hello-blockchain-world/testutil/nullify"
	"github.com/Zireael26/hello-blockchain-world/x/helloblockchainworld"
	"github.com/Zireael26/hello-blockchain-world/x/helloblockchainworld/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloblockchainworldKeeper(t)
	helloblockchainworld.InitGenesis(ctx, *k, genesisState)
	got := helloblockchainworld.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
