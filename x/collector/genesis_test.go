package collector_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	app "github.com/redactedfury/petri/app"
	"github.com/redactedfury/petri/x/collector"
	"github.com/redactedfury/petri/x/collector/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	petriApp := app.Setup(false)
	ctx := petriApp.BaseApp.NewContext(false, tmproto.Header{})
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k := petriApp.CollectorKeeper
	collector.InitGenesis(ctx, k, &genesisState)
	got := collector.ExportGenesis(ctx, k)
	require.NotNil(t, got)
}
