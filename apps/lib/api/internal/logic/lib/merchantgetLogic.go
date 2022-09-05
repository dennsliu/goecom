package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantgetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantgetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantgetLogic {
	return &MerchantgetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantgetLogic) Merchantget(req *types.MerchantGetReq) (resp *types.MerchantReply, err error) {
	// todo: add your logic here and delete this line

	return
}
