package keeper

import (
    "encoding/binary"

    "carreg/x/carreg/types"

    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendIvc(ctx sdk.Context, post types.Ivc) uint64 {
    count := k.GetIvcCount(ctx)
    post.Id = count
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.IvcKey))
    appendedValue := k.cdc.MustMarshal(&post)
    store.Set(GetIvcIDBytes(post.Id), appendedValue)
    k.SetIvcCount(ctx, count+1)
    return count
}

func (k Keeper) GetIvcCount(ctx sdk.Context) uint64 {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, []byte{})
    byteKey := types.KeyPrefix(types.IvcCountKey)
    bz := store.Get(byteKey)
    if bz == nil {
        return 0
    }
    return binary.BigEndian.Uint64(bz)
}

func GetIvcIDBytes(id uint64) []byte {
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, id)
    return bz
}

func (k Keeper) SetIvcCount(ctx sdk.Context, count uint64) {
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, []byte{})
    byteKey := types.KeyPrefix(types.IvcCountKey)
    bz := make([]byte, 8)
    binary.BigEndian.PutUint64(bz, count)
    store.Set(byteKey, bz)
}