package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

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

	return
}
