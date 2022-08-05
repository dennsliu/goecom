package logic

import (
	"context"
	"strings"
	"time"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantaddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantaddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantaddLogic {
	return &MerchantaddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantaddLogic) Merchantadd(req *types.MerchantAddReq) (resp *types.Merchant, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Name)) == 0 {
		logx.Errorf("Merchantuseradd err null name:%s", req.Name)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantadd err null name:%s", req.Name)
	}
	merchantModel := new(model.Merchant)
	merchantModel.Name = req.Name
	merchantModel.Status = req.Status
	merchantModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	merchantModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	insertResult, err := l.svcCtx.MerchantModel.Insert(l.ctx, merchantModel)
	if err != nil {
		logx.Errorf("Merchantadd insert error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantadd insert err name:%s, err:%v", req.Name, err)
	}
	lastId, err := insertResult.LastInsertId()

	merchantInfo, err := l.svcCtx.MerchantModel.FindOne(l.ctx, lastId)
	if err != nil {
		logx.Errorf("Merchantuseradd find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuseradd find err id:%d, err:%v", lastId, err)
	}
	return &types.Merchant{
		Id:        merchantInfo.Id,
		Name:      merchantInfo.Name,
		Status:    merchantInfo.Status,
		CreatedAt: merchantInfo.CreatedAt,
		UpdatedAt: merchantInfo.UpdatedAt,
	}, nil
}
