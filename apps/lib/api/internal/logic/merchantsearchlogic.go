package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsearchLogic {
	return &MerchantsearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsearchLogic) Merchantsearch(req *types.MerchantSearchReq) (resp *types.MerchantSearchReply, err error) {
	// todo: add your logic here and delete this line

	return
}
