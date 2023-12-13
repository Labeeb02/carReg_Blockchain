package keeper

import (
	"context"

	"carreg/x/carreg/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateIvc(goCtx context.Context, msg *types.MsgCreateIvc) (*types.MsgCreateIvcResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    var post = types.Ivc{
        Creator: msg.Creator,
        VehicleNo:   msg.VehicleNo,
        Owner:    msg.Owner,
		RaSign: msg.RaSign,
    }
    id := k.AppendIvc(
        ctx,
        post,
    )
    return &types.MsgCreateIvcResponse{
        Id: id,
    }, nil
}