package carreg_test

import (
	"testing"

	keepertest "carreg/testutil/keeper"
	"carreg/testutil/nullify"
	"carreg/x/carreg/module"
	"carreg/x/carreg/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarregKeeper(t)
	carreg.InitGenesis(ctx, k, genesisState)
	got := carreg.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
