package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

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

	return
}
