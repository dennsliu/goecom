package logic

import (
	"context"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantUserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantUserRegisterLogic {
	return &MerchantUserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MerchantUserRegisterLogic) MerchantUserRegister(in *lib.MerchantUserRegisterReq) (*lib.MerchantUserLoginReply, error) {
	// todo: add your logic here and delete this line

	return &lib.MerchantUserLoginReply{}, nil
}
