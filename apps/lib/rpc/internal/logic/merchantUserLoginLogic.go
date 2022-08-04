package logic

import (
	"context"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantUserLoginLogic {
	return &MerchantUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MerchantUserLoginLogic) MerchantUserLogin(in *lib.MerchantUserLoginReq) (*lib.MerchantUserLoginReply, error) {
	// todo: add your logic here and delete this line

	return &lib.MerchantUserLoginReply{}, nil
}
