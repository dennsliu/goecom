package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LanguagesdeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLanguagesdeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LanguagesdeleteLogic {
	return &LanguagesdeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LanguagesdeleteLogic) Languagesdelete(req *types.LanguagesDeleteReq) (resp *types.LanguagesDeleteReply, err error) {
	// todo: add your logic here and delete this line

	return
}
