package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/tejiriaustin/hello-web3/testutil/keeper"
	"github.com/tejiriaustin/hello-web3/x/helloweb3/keeper"
	"github.com/tejiriaustin/hello-web3/x/helloweb3/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Helloweb3Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
