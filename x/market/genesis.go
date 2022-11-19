package market

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/redactedfury/sxfury/x/market/keeper"
	"github.com/redactedfury/sxfury/x/market/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {
	for _, item := range state.TimeWeightedAverage {
		k.SetTwa(ctx, item)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(k.GetAllTwa(ctx))
}
