package simulation

import (
	"github.com/Carina-hackatom/nova/x/gal/keeper"
	"github.com/Carina-hackatom/nova/x/gal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// WeightedOperations returns all the operations from the module with their respective seights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, k keeper.Keeper, ak types.AccountKeeper,
) simulation.WeightedOperations {
	// TODO : implements this!
	return nil
}
