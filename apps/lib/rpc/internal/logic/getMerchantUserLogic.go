package logic

import (
	"context"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantUserLogic {
	return &GetMerchantUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantUserLogic) GetMerchantUser(in *lib.GetMerchantUserReq) (*lib.GetMerchantUserReply, error) {
	// todo: add your logic here and delete this line

	return &lib.GetMerchantUserReply{}, nil
}
