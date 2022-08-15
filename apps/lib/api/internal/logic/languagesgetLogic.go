package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

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

	return
}
