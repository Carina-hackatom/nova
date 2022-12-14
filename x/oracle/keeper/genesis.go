package keeper

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/Carina-hackatom/nova/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	for i := range genState.States {
		if err := k.UpdateChainState(ctx, &genState.States[i]); err != nil {
			panic(fmt.Errorf("failed to initialize genesis state at %s, err: %v", types.ModuleName, err))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	result := &types.GenesisState{
		Params: k.GetParams(ctx),
		States: []types.ChainInfo{},
	}

	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer func() {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
		}
	}()

	for ; iterator.Valid(); iterator.Next() {
		value := types.ChainInfo{}
		if err := proto.Unmarshal(iterator.Value(), &value); err != nil {
			panic(fmt.Errorf("unable to unmarshal chain state: %v", err))
		}

		result.States = append(result.States, types.ChainInfo{
			Coin:            value.Coin,
			OperatorAddress: value.OperatorAddress,
			LastBlockHeight: value.LastBlockHeight,
			Decimal:         value.Decimal,
		})
	}

	return result
}
