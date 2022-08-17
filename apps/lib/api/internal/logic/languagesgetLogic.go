package logic

import (
	"context"
	"fmt"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LanguagesgetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLanguagesgetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LanguagesgetLogic {
	return &LanguagesgetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LanguagesgetLogic) Languagesget(req *types.LanguagesGetReq) (resp *types.LanguagesReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Print("--------------Languagesget-------------")
	if req.Id == 0 {
		return nil, errors.Wrapf(ErrUsernamePwdNullError, "Languagesget err null id:%d", req.Id)
	}
	languageInfo, err := l.svcCtx.LanguagesModel.FindOne(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Languagesget find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Languagesget find err id:%d, err:%v", req.Id, err)
	}
	return &types.LanguagesReply{
		Code:         200,
		Msg:          "get language info successfully",
		Id:           languageInfo.Id,
		Name:         languageInfo.Name,
		Languagecode: languageInfo.Code,
		Image:        languageInfo.Image,
		Directory:    languageInfo.Directory,
		CreatedAt:    languageInfo.CreatedAt,
		UpdatedAt:    languageInfo.UpdatedAt,
	}, nil
}
