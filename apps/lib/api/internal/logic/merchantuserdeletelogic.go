package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuserdeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuserdeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuserdeleteLogic {
	return &MerchantuserdeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuserdeleteLogic) Merchantuserdelete(req *types.MerchantUserDeleteReq) (resp *types.MerchantUserDeleteReply, err error) {
	// todo: add your logic here and delete this line

	return
}
