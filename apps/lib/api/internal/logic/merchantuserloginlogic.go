package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuserloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuserloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuserloginLogic {
	return &MerchantuserloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuserloginLogic) Merchantuserlogin(req *types.MerchantUserLoginReq) (resp *types.Merchant, err error) {
	// todo: add your logic here and delete this line

	return
}
