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

type LanguagesaddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLanguagesaddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LanguagesaddLogic {
	return &LanguagesaddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LanguagesaddLogic) Languagesadd(req *types.LanguagesAddReq) (resp *types.LanguagesReply, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Name)) == 0 {
		logx.Errorf("Merchantuseradd err null name:%s", req.Name)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantadd err null name:%s", req.Name)
	}
	languagesModel := new(model.Languages)
	languagesModel.Name = req.Name
	languagesModel.Code = req.Languagecode
	languagesModel.Image = req.Image
	languagesModel.Directory = req.Directory
	languagesModel.Image = req.Image
	languagesModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	languagesModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	insertResult, err := l.svcCtx.LanguagesModel.Insert(l.ctx, languagesModel)
	if err != nil {
		logx.Errorf("Languagesadd insert error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Languagesadd insert err name:%s, err:%v", req.Name, err)
	}
	lastId, err := insertResult.LastInsertId()

	LanguageInfo, err := l.svcCtx.LanguagesModel.FindOne(l.ctx, lastId)
	if err != nil {
		logx.Errorf("Languagesadd find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Languagesadd find err id:%d, err:%v", lastId, err)
	}
	return &types.LanguagesReply{
		Id:           LanguageInfo.Id,
		Name:         LanguageInfo.Name,
		Languagecode: LanguageInfo.Code,
		Image:        LanguageInfo.Image,
		Directory:    LanguageInfo.Directory,
		Order:        LanguageInfo.Order,
		CreatedAt:    LanguageInfo.CreatedAt,
		UpdatedAt:    LanguageInfo.UpdatedAt,
	}, nil
}
