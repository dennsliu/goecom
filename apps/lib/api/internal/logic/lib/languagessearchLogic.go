package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LanguagessearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLanguagessearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LanguagessearchLogic {
	return &LanguagessearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LanguagessearchLogic) Languagessearch(req *types.LanguagesSearchReq) (resp *types.LanguagesSearchReply, err error) {
	// todo: add your logic here and delete this line

	return
}
