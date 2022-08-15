package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
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
	if req.Id == 0 {
		logx.Errorf("Languagesdelete err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Languagesdelete err null id:%s", req.Id)
	}
	err = l.svcCtx.LanguagesModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Languagesdelete delete error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Languagesdelete delete err id:%s, err:%v", req.Id, err)
	}
	return &types.LanguagesDeleteReply{
		Code: 200,
		Msg:  "Delete language successfully",
		Id:   req.Id,
	}, nil
}
