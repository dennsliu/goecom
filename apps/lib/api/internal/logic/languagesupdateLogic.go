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

type LanguagesupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLanguagesupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LanguagesupdateLogic {
	return &LanguagesupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LanguagesupdateLogic) Languagesupdate(req *types.LanguagesUpdateReq) (resp *types.LanguagesReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id == 0 {
		logx.Errorf("Languagesupdate err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Languagesupdate err null id:%s", req.Id)
	}
	if len(strings.TrimSpace(req.Name)) == 0 {
		logx.Errorf("Languagesupdate err null name:%s", req.Name)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Languagesupdate err null name:%s", req.Name)
	}
	originalData, err := l.svcCtx.LanguagesModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Languagesupdate FindOne error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Languagesupdate findOne err id:%d, err:%v", req.Id, err)
	}
	originalData.Name = req.Name
	originalData.Code = req.Languagecode
	originalData.Image = req.Image
	originalData.Directory = req.Directory
	originalData.Order = req.Order
	originalData.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err = l.svcCtx.LanguagesModel.Update(l.ctx, originalData)
	if err != nil {
		logx.Errorf("Languagesupdate update error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Languagesupdate update err id:%s, err:%v", req.Id, err)
	}
	return &types.LanguagesReply{
		Code:         200,
		Msg:          "Update language successfully",
		Id:           originalData.Id,
		Name:         originalData.Name,
		Languagecode: originalData.Code,
		Image:        originalData.Image,
		Directory:    originalData.Directory,
		Order:        originalData.Order,
		CreatedAt:    originalData.CreatedAt,
		UpdatedAt:    originalData.UpdatedAt,
	}, nil
}
