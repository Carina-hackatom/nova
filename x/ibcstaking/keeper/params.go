package keeper

import (
	"github.com/Carina-hackatom/nova/x/ibcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParams sets the ibcstaking module's parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetParams gets the ibcstaking module's parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}
