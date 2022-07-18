package keeper

import (
	"context"

	cosm "github.com/cosmos/cosmos-sdk/types"
	"github.com/cytruslabs/zestchain/x/zestchain/types"
)

func (k msgServer) CreatePromo(goCtx context.Context, msg *types.MsgCreatePromo) (*types.MsgCreatePromoResponse, error) {
	ctx := cosm.UnwrapSDKContext(goCtx)

	promoCount, found := k.Keeper.GetPromoCount(ctx)
	if !found {
		panic("PromoCount not found")
	}
	newIndex := strconv.FormatUint(promoCount.Total, 10)
	//pot, _ := strconv.ParseUint(msg.Pot, 10, 64)
	ad := types.Ad{
		Index:   newIndex,
		Title:   msg.Title,
		Pot:     msg.Pot,
		Url:     msg.Url,
		Msg:     msg.Msg,
		Tags:    msg.Tags,
		Prefs:   msg.Prefs,
		Creator: msg.Creator,
	}

	k.Keeper.SetPromo(ctx, promo)
	promoCount.Total++
	k.Keeper.SetPromoCount(ctx, promoCount)

	return &types.MsgCreatePromoResponse{
		Total: newIndex,
	}, nil
}
