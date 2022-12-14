package gal

import (
	"github.com/Carina-hackatom/nova/x/gal/keeper"
	"github.com/Carina-hackatom/nova/x/gal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the gal module's state from a given genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the gal module's genesis state.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesisState()
}
