package logic

import (
	"context"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuserdeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuserdeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuserdeleteLogic {
	return &MerchantuserdeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuserdeleteLogic) Merchantuserdelete(req *types.MerchantUserDeleteReq) (resp *types.MerchantUserDeleteReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id == 0 {
		logx.Errorf("Languagesdelete err null id:%s", req.Id)
		return nil, errors.Wrapf(ErrMerchantNameNullError, "Merchantuserdelete err null id:%s", req.Id)
	}
	err = l.svcCtx.MerchantUserModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Errorf("Merchantuserdelete delete error: %v", err)
		return nil, errors.Wrapf(ErrMerchantAddError, "Merchantuserdelete delete err id:%s, err:%v", req.Id, err)
	}
	return &types.MerchantUserDeleteReply{
		Code: 200,
		Msg:  "Delete merchant user successfully",
		Id:   req.Id,
	}, nil
}
