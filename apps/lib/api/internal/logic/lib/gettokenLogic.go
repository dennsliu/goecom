package lib

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GettokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGettokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GettokenLogic {
	return &GettokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GettokenLogic) Gettoken(req *types.GetTokenReq) (resp *types.GetTokenReply, err error) {
	// todo: add your logic here and delete this line

	return
}
