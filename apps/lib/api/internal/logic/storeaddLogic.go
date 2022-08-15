package logic

import (
	"context"
	"time"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type StoreaddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreaddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreaddLogic {
	return &StoreaddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreaddLogic) Storeadd(req *types.StoreAddReq) (resp *types.StoreReply, err error) {
	// todo: add your logic here and delete this line
	storeModel := new(model.Store)
	storeModel.MerchantId = req.Merchantid
	storeModel.Status = req.Status
	storeModel.Order = req.Order
	storeModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	storeModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	insertResult, err := l.svcCtx.StoreModel.Insert(l.ctx, storeModel)
	lastId, err := insertResult.LastInsertId()
	if len(req.StoreLaguage) > 0 {

	}
	storeInfo, err := l.svcCtx.StoreModel.FindOne(l.ctx, lastId)
	if err != nil {
		logx.Errorf("Storeadd find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Storeadd find err id:%d, err:%v", lastId, err)
	}
	return &types.StoreReply{
		Code:       200,
		Msg:        "add store successfully",
		Id:         storeInfo.Id,
		Order:      storeInfo.Order,
		Merchantid: storeInfo.MerchantId,
		CreatedAt:  storeInfo.CreatedAt,
		UpdatedAt:  storeInfo.UpdatedAt,
	}, nil
}
