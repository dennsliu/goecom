package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type StoregetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoregetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoregetLogic {
	return &StoregetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoregetLogic) Storeget(req *types.StoreGetReq) (resp *types.StoreReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Print("--------------Storeget-------------")
	if req.Id == 0 {
		return nil, errors.Wrapf(ErrUsernamePwdNullError, "Storeget err null id:%d", req.Id)
	}
	storeInfo, err := l.svcCtx.StoreModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Storeget find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Storeget find err id:%d, err:%v", req.Id, err)
	}
	return &types.StoreReply{
		Code:       200,
		Msg:        "get store info successfully",
		Id:         storeInfo.Id,
		MerchantId: storeInfo.MerchantId,
		Order:      storeInfo.Order,
		CreatedAt:  storeInfo.CreatedAt,
		UpdatedAt:  storeInfo.UpdatedAt,
	}, nil
}
