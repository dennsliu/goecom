package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantaddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantaddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantaddLogic {
	return &MerchantaddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantaddLogic) Merchantadd(req *types.MerchantAddReq) (resp *types.MerchantReply, err error) {
	// todo: add your logic here and delete this line

	return
}
