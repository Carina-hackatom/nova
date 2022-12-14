package keeper

import (
	"fmt"

	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"

	"github.com/Carina-hackatom/nova/x/ibcstaking/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterZone
func (k Keeper) RegisterZone(ctx sdk.Context, zone *types.RegisteredZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := k.cdc.MustMarshal(zone)
	store.Set([]byte(zone.ZoneId), bz)
}

// GetRegisteredZone
func (k Keeper) GetRegisteredZone(ctx sdk.Context, zoneId string) (types.RegisteredZone, bool) {
	zone := types.RegisteredZone{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	bz := store.Get([]byte(zoneId))

	if len(bz) == 0 {
		return zone, false
	}

	k.cdc.MustUnmarshal(bz, &zone)
	return zone, true
}

// DeleteRegisteredZone delete zone info
func (k Keeper) DeleteRegisteredZone(ctx sdk.Context, zoneId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	ctx.Logger().Error(fmt.Sprintf("Removing chain: %s", zoneId))
	store.Delete([]byte(zoneId))
}

// IterateRegisteredZones iterate through zones
func (k Keeper) IterateRegisteredZones(ctx sdk.Context, fn func(index int64, zoneInfo types.RegisteredZone) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixZone)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer func(iterator sdk.Iterator) {
		err := iterator.Close()
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("unexpected iterator closed: %s", err))
			return
		}
	}(iterator)

	i := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		zone := types.RegisteredZone{}
		k.cdc.MustUnmarshal(iterator.Value(), &zone)

		stop := fn(i, zone)
		if stop {
			break
		}
		i++
	}
}

func (k Keeper) GetRegisteredZoneForValidatorAddr(ctx sdk.Context, validatorAddr string) *types.RegisteredZone {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.ValidatorAddress == validatorAddr {
			zone = &zoneInfo
			return true
		}
		return false
	})
	return zone
}

func (k Keeper) GetZoneForDenom(ctx sdk.Context, denom string) *types.RegisteredZone {
	var zone *types.RegisteredZone

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == denom {
			zone = &zoneInfo
			return true
		}
		return false
	})

	return zone
}

func (k Keeper) GetsnDenomForBaseDenom(ctx sdk.Context, baseDenom string) string {
	var snDenom string

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.BaseDenom == baseDenom {
			snDenom = zoneInfo.SnDenom
			return true
		}
		return false
	})

	return snDenom
}

func (k Keeper) GetBaseDenomForSnDenom(ctx sdk.Context, snDenom string) string {
	var baseDenom string

	k.IterateRegisteredZones(ctx, func(_ int64, zoneInfo types.RegisteredZone) (stop bool) {
		if zoneInfo.SnDenom == snDenom {
			baseDenom = zoneInfo.BaseDenom
			return true
		}
		return false
	})
	return baseDenom
}

func (k Keeper) GetIBCHashDenom(ctx sdk.Context, portId, chanId, baseDenom string) string {
	var path string

	if portId == "" || chanId == "" {
		path = ""
	} else {
		path = portId + "/" + chanId
	}

	denomTrace := transfertypes.DenomTrace{
		Path:      path,
		BaseDenom: baseDenom,
	}

	denomHash := denomTrace.IBCDenom()

	return denomHash
}
