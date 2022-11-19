package wasm

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	assetkeeper "github.com/redactedfury/sxfury/x/asset/keeper"
	auctionKeeper "github.com/redactedfury/sxfury/x/auction/keeper"
	collectorKeeper "github.com/redactedfury/sxfury/x/collector/keeper"
	esmKeeper "github.com/redactedfury/sxfury/x/esm/keeper"
	lendKeeper "github.com/redactedfury/sxfury/x/lend/keeper"
	liquidationKeeper "github.com/redactedfury/sxfury/x/liquidation/keeper"
	liquidityKeeper "github.com/redactedfury/sxfury/x/liquidity/keeper"
	lockerkeeper "github.com/redactedfury/sxfury/x/locker/keeper"
	rewardsKeeper "github.com/redactedfury/sxfury/x/rewards/keeper"
	tokenMintkeeper "github.com/redactedfury/sxfury/x/tokenmint/keeper"
	vaultKeeper "github.com/redactedfury/sxfury/x/vault/keeper"
)

func RegisterCustomPlugins(
	locker *lockerkeeper.Keeper,
	tokenMint *tokenMintkeeper.Keeper,
	asset *assetkeeper.Keeper,
	rewards *rewardsKeeper.Keeper,
	collector *collectorKeeper.Keeper,
	liquidation *liquidationKeeper.Keeper,
	auction *auctionKeeper.Keeper,
	esm *esmKeeper.Keeper,
	vault *vaultKeeper.Keeper,
	lend *lendKeeper.Keeper,
	liquidity *liquidityKeeper.Keeper,
) []wasmkeeper.Option {
	petriQueryPlugin := NewQueryPlugin(asset, locker, tokenMint, rewards, collector, liquidation, esm, vault, lend, liquidity)

	appDataQueryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(petriQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(*locker, *rewards, *asset, *collector, *liquidation, *auction, *tokenMint, *esm, *vault),
	)

	return []wasm.Option{
		appDataQueryPluginOpt,
		messengerDecoratorOpt,
	}
}
