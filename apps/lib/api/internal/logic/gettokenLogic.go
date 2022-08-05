package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/rpc/lib"

	"github.com/pkg/errors"
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

func (l *GettokenLogic) Gettoken(req *types.GetTokenReq) (*types.GetTokenReply, error) {
	// todo: add your logic here and delete this line
	tokenResp, err := l.svcCtx.LibRpc.GenerateToken(l.ctx, &lib.GenerateTokenReq{
		UserId: 0,
	})
	if err != nil {
		logx.Errorf("get token error: %v", err)
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", 0, err)
	}
	var resp types.GetTokenReply
	//_ = copier.Copy(&resp, tokenResp)
	resp.AccessExpire = tokenResp.AccessExpire
	resp.AccessToken = tokenResp.AccessToken
	resp.RefreshAfter = tokenResp.RefreshAfter
	return &resp, nil
}
