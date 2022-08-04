package logic

import (
	"context"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantUsersLogic {
	return &GetMerchantUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantUsersLogic) GetMerchantUsers(in *lib.GetMerchantUsersReq) (*lib.GetMerchantUsersReply, error) {
	// todo: add your logic here and delete this line

	return &lib.GetMerchantUsersReply{}, nil
}
