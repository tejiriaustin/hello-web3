package helloweb3_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tejiriaustin/hello-web3/testutil/keeper"
	"github.com/tejiriaustin/hello-web3/testutil/nullify"
	"github.com/tejiriaustin/hello-web3/x/helloweb3"
	"github.com/tejiriaustin/hello-web3/x/helloweb3/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Helloweb3Keeper(t)
	helloweb3.InitGenesis(ctx, *k, genesisState)
	got := helloweb3.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
