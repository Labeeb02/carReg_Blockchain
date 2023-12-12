package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "carreg/testutil/keeper"
	"carreg/x/carreg/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CarregKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
