package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantgetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantgetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantgetLogic {
	return &MerchantgetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantgetLogic) Merchantget(req *types.MerchantGetReq) (resp *types.MerchantReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Print("--------------Merchantget-------------")
	if req.Id == 0 {
		return nil, errors.Wrapf(ErrUsernamePwdNullError, "Merchantget err null id:%d", req.Id)
	}
	merchantInfo, err := l.svcCtx.MerchantModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Merchantget find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantget find err id:%d, err:%v", req.Id, err)
	}
	return &types.MerchantReply{
		Code:      200,
		Msg:       "get merchant info successfully",
		Id:        merchantInfo.Id,
		Name:      merchantInfo.Name,
		Status:    merchantInfo.Status,
		CreatedAt: merchantInfo.CreatedAt,
		UpdatedAt: merchantInfo.UpdatedAt,
	}, nil
}
