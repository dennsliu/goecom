package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantupdateLogic {
	return &MerchantupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantupdateLogic) Merchantupdate(req *types.MerchantUpdateReq) (resp *types.Merchant, err error) {
	// todo: add your logic here and delete this line

	return
}
