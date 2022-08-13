package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantdeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantdeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantdeleteLogic {
	return &MerchantdeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantdeleteLogic) Merchantdelete(req *types.MerchantDeleteReq) (resp *types.MerchantDeleteReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id == 0 {
		logx.Errorf("Merchantdelete err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantdelete err null id:%s", req.Id)
	}
	err = l.svcCtx.MerchantModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Merchantdelete delete error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantdelete delete err id:%s, err:%v", req.Id, err)
	}
	return &types.MerchantDeleteReply{
		Code: 200,
		Msg:  "Delete merchant successfully",
		Id:   req.Id,
	}, nil
}
