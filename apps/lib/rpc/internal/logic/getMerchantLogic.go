package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/rpc/internal/svc"
	"goecom/apps/lib/rpc/types/lib"
	"goecom/pkg/xerr"

	"github.com/pkg/errors"
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
	fmt.Print("--------------Merchantget rpc-------------id:%d", in.Id)
	// if in.Id == 0 {
	// 	return nil, errors.Wrapf(ErrUsernamePwdNullError, "Merchantget err null id:%d", in.Id)
	// }
	merchantInfo, err := l.svcCtx.MerchantModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.Errorf("Merchantgetrpc find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantgetrpc find err id:%d, err:%v", in.Id, err)
	}
	return &lib.GetMerchantReply{
		Id:        merchantInfo.Id,
		Name:      merchantInfo.Name,
		Status:    merchantInfo.Status,
		Createdat: merchantInfo.CreatedAt,
		Updatedat: merchantInfo.UpdatedAt,
	}, nil
}
