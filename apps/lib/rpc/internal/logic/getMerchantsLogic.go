package logic

import (
	"context"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantsLogic {
	return &GetMerchantsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantsLogic) GetMerchants(in *lib.GetMerchantsReq) (*lib.GetMerchantsReply, error) {
	// todo: add your logic here and delete this line

	return &lib.GetMerchantsReply{}, nil
}
