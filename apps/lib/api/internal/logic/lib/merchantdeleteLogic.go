package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantdeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantdeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantdeleteLogic {
	return &MerchantdeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantdeleteLogic) Merchantdelete(req *types.MerchantDeleteReq) (resp *types.MerchantDeleteReply, err error) {
	// todo: add your logic here and delete this line

	return
}
