package logic

import (
	"context"
	"fmt"
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
	storeModel.MerchantId = req.MerchantId
	storeModel.Status = req.Status
	storeModel.Order = req.Order
	storeModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	storeModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	insertResult, err := l.svcCtx.StoreModel.Insert(l.ctx, storeModel)
	lastId, err := insertResult.LastInsertId()
	if len(req.StoreLaguage) > 0 {
		for _, storeLanguage := range req.StoreLaguage {
			fmt.Printf("------------result------storeLanguage:")
			fmt.Print(storeLanguage)
			storeLanguageModel := new(model.StoreLanguage)
			storeLanguageModel.Storeid = lastId
			storeLanguageModel.Name = storeLanguage.Name
			storeLanguageModel.Keyword = storeLanguage.Keyword
			storeLanguageModel.Description = storeLanguage.Description
			storeLanguageModel.Laguageid = storeLanguage.Laguageid
			storeLanguageModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			storeLanguageModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			insertLanguageResult, err := l.svcCtx.StoreModel.InsertLanguage(l.ctx, storeLanguageModel)
			fmt.Print(insertLanguageResult)
			fmt.Print(err)
		}
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
		MerchantId: storeInfo.MerchantId,
		CreatedAt:  storeInfo.CreatedAt,
		UpdatedAt:  storeInfo.UpdatedAt,
	}, nil
}
