package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantusergetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantusergetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantusergetLogic {
	return &MerchantusergetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantusergetLogic) Merchantuserget(req *types.MerchantUserGetReq) (resp *types.MerchantUser, err error) {
	// todo: add your logic here and delete this line

	return
}
