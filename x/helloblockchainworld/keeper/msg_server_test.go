package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Zireael26/hello-blockchain-world/testutil/keeper"
	"github.com/Zireael26/hello-blockchain-world/x/helloblockchainworld/keeper"
	"github.com/Zireael26/hello-blockchain-world/x/helloblockchainworld/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HelloblockchainworldKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
