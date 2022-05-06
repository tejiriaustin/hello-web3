package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tejiriaustin/hello-web3/testutil/keeper"
	"github.com/tejiriaustin/hello-web3/x/helloweb3/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Helloweb3Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
