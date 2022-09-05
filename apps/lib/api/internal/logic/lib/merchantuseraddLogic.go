package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuseraddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuseraddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuseraddLogic {
	return &MerchantuseraddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuseraddLogic) Merchantuseradd(req *types.MerchantUserAddReq) (resp *types.MerchantUserReply, err error) {
	// todo: add your logic here and delete this line

	return
}
