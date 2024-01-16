package blogapp_test

import (
	"testing"

	keepertest "blogapp/testutil/keeper"
	"blogapp/testutil/nullify"
	"blogapp/x/blogapp/module"
	"blogapp/x/blogapp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogappKeeper(t)
	blogapp.InitGenesis(ctx, k, genesisState)
	got := blogapp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
