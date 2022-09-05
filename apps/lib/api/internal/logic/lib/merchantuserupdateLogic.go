package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuserupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuserupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuserupdateLogic {
	return &MerchantuserupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuserupdateLogic) Merchantuserupdate(req *types.MerchantUserUpdateReq) (resp *types.MerchantUserReply, err error) {
	// todo: add your logic here and delete this line

	return
}
