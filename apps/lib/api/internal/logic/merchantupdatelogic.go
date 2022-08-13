package logic

import (
	"context"
	"strings"
	"time"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantupdateLogic {
	return &MerchantupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantupdateLogic) Merchantupdate(req *types.MerchantUpdateReq) (resp *types.MerchantReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id == 0 {
		logx.Errorf("Merchantupdate err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantupdate err null id:%s", req.Id)
	}
	if len(strings.TrimSpace(req.Name)) == 0 {
		logx.Errorf("Merchantupdate err null name:%s", req.Name)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantupdate err null name:%s", req.Name)
	}
	originalData, err := l.svcCtx.MerchantModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Merchantupdate FindOne error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantupdate findOne err id:%d, err:%v", req.Id, err)
	}
	originalData.Name = req.Name
	originalData.Status = req.Status
	originalData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = l.svcCtx.MerchantModel.Update(l.ctx, originalData)
	if err != nil {
		logx.Errorf("Merchantupdate update error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantupdate update err id:%s, err:%v", req.Id, err)
	}
	return &types.MerchantReply{
		Code:      200,
		Msg:       "Update merchant successfully",
		Id:        originalData.Id,
		Name:      originalData.Name,
		Status:    originalData.Status,
		CreatedAt: originalData.CreatedAt,
		UpdatedAt: originalData.UpdatedAt,
	}, nil
}
