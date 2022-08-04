package logic

import (
	"context"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"
	"goecom/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewGetMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantLogic {
	return &GetMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantLogic) GetMerchant(in *lib.GetMerchantReq) (*lib.GetMerchantReply, error) {
	// todo: add your logic here and delete this line

	return &lib.GetMerchantReply{}, nil
}
