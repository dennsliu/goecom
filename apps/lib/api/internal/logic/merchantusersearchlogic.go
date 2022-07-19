package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantusersearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantusersearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantusersearchLogic {
	return &MerchantusersearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantusersearchLogic) Merchantusersearch(req *types.MerchantSearchReq) (resp *types.MerchantUserSearchReply, err error) {
	// todo: add your logic here and delete this line

	return
}
