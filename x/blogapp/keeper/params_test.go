package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "blogapp/testutil/keeper"
	"blogapp/x/blogapp/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BlogappKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
