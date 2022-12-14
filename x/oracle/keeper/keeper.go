package keeper

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/Carina-hackatom/nova/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc: cdc, storeKey: key, paramSpace: paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) UpdateChainState(ctx sdk.Context, chainInfo *types.ChainInfo) error {
	if !k.IsValidOperator(ctx, chainInfo.OperatorAddress) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, chainInfo.OperatorAddress)
	}

	store := ctx.KVStore(k.storeKey)
	bz, err := chainInfo.Marshal()
	if err != nil {
		return err
	}

	store.Set([]byte(chainInfo.Coin.Denom), bz)
	return nil
}

func (k Keeper) GetChainState(ctx sdk.Context, chainDenom string) (*types.ChainInfo, error) {
	chainInfo := types.ChainInfo{}
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(chainDenom)) {
		return nil, sdkerrors.Wrap(types.ErrNoSupportChain, fmt.Sprintf("chain %s", chainDenom))
	}

	data := store.Get([]byte(chainDenom))
	if err := proto.Unmarshal(data, &chainInfo); err != nil {
		return nil, err
	}

	return &chainInfo, nil
}

func (k Keeper) IsValidOperator(ctx sdk.Context, operatorAddress string) bool {
	params := k.GetParams(ctx)

	for i := range params.OracleOperators {
		if params.OracleOperators[i] == operatorAddress {
			return true
		}
	}

	return false
}

func (k Keeper) GetOracleVersion(ctx sdk.Context, chainDenom string) (int64, error) {
	chainInfo, err := k.GetChainState(ctx, chainDenom)
	if err != nil {
		return 0, err
	}

	return chainInfo.LastBlockHeight, nil
}
