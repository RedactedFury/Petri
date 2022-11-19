package tokenmint_test

// import (
// 	"testing"

// 	keepertest "github.com/redactedfury/sxfury/testutil/keeper"
// 	"github.com/redactedfury/sxfury/testutil/nullify"
// 	"github.com/redactedfury/sxfury/x/tokenmint"
// 	"github.com/redactedfury/sxfury/x/tokenmint/types"
// 	"github.com/stretchr/testify/require"
// )

// func TestGenesis(t *testing.T) {
// 	genesisState := types.GenesisState{
// 		Params:	types.DefaultParams(),

// 	}

// 	k, ctx := keepertest.TokenmintKeeper(t)
// 	tokenmint.InitGenesis(ctx, *k, genesisState)
// 	got := tokenmint.ExportGenesis(ctx, *k)
// 	require.NotNil(t, got)

// 	nullify.Fill(&genesisState)
// 	nullify.Fill(got)

// }
