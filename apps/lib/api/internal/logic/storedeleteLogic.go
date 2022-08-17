package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type StoredeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoredeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoredeleteLogic {
	return &StoredeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoredeleteLogic) Storedelete(req *types.StoreDeleteReq) (resp *types.StoreDeleteReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id == 0 {
		logx.Errorf("Storedelete err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantuserdelete err null id:%s", req.Id)
	}
	err = l.svcCtx.MerchantModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Storedelete delete error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Storedelete delete err id:%s, err:%v", req.Id, err)
	}
	return &types.StoreDeleteReply{
		Code: 200,
		Msg:  "Delete store successfully",
		Id:   req.Id,
	}, nil
}
