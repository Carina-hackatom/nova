package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Carina-hackatom/nova/x/ibcstaking/types"
)

type Keeper struct {
	cdc codec.Codec

	storeKey            sdk.StoreKey
	authKeeper          types.AccountKeeper
	scopedKeeper        capabilitykeeper.ScopedKeeper
	icaControllerKeeper icacontrollerkeeper.Keeper
	hooks               types.ICAHooks
	paramSpace          paramtypes.Subspace
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, ak types.AccountKeeper, iaKeeper icacontrollerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper, paramStore paramtypes.Subspace) Keeper {
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	if !paramStore.HasKeyTable() {
		paramStore = paramStore.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:                 cdc,
		storeKey:            storeKey,
		authKeeper:          ak,
		scopedKeeper:        scopedKeeper,
		paramSpace:          paramStore,
		icaControllerKeeper: iaKeeper,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s-%s", host.ModuleName, types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// SetHooks set the epoch hooks
func (k *Keeper) SetHooks(eh types.ICAHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set ICA hooks twice")
	}

	k.hooks = eh

	return k
}

func (k Keeper) IsValidDaoModifier(ctx sdk.Context, daoModifier string) bool {
	params := k.GetParams(ctx)
	for i := range params.DaoModifiers {
		if params.DaoModifiers[i] == daoModifier {
			return true
		}
	}
	return false
}
