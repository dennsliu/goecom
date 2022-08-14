package logic

import (
	"context"
	"time"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuserupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuserupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuserupdateLogic {
	return &MerchantuserupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuserupdateLogic) Merchantuserupdate(req *types.MerchantUserUpdateReq) (resp *types.MerchantUserReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id == 0 {
		logx.Errorf("Merchantuserupdate err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantupdate err null id:%s", req.Id)
	}

	originalData, err := l.svcCtx.MerchantUserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Merchantuserupdate FindOne error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantuserupdate findOne err id:%d, err:%v", req.Id, err)
	}
	originalData.Nickname = req.Nickname
	originalData.Telephone = req.Telephone
	originalData.Mobliephone = req.Mobliephone
	originalData.Status = req.Status
	originalData.MerchantId = req.MerchantId
	originalData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = l.svcCtx.MerchantUserModel.Update(l.ctx, originalData)
	if err != nil {
		logx.Errorf("Merchantupdate update error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantupdate update err id:%s, err:%v", req.Id, err)
	}
	return &types.MerchantUserReply{
		Code:        200,
		Msg:         "Update merchant user successfully",
		Id:          originalData.Id,
		Nickname:    originalData.Nickname,
		Username:    originalData.Username,
		Telephone:   originalData.Telephone,
		Mobliephone: originalData.Mobliephone,
		Status:      originalData.Status,
		MerchantId:  originalData.MerchantId,
		CreatedAt:   originalData.CreatedAt,
		UpdatedAt:   originalData.UpdatedAt,
	}, nil
}
