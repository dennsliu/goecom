package logic

import (
	"context"
	"goecom/pkg/xerr"
	"strings"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuserloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuserloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuserloginLogic {
	return &MerchantuserloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantuserloginLogic) Merchantuserlogin(req *types.MerchantUserLoginReq) (resp *types.MerchantUser, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("Username and Password are required"), "req: %+v,api err:%+v", req, err)
	}
	userInfo, err := l.svcCtx.MerchantUserModel.UserLogin(l.ctx, req.UserName)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.Wrapf(xerr.NewErrMsg("user not found"), "req: %+v,api err:%+v", req, err)
	default:
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, errors.Wrapf(xerr.NewErrMsg("password mismatched"), "req: %+v,api err:%+v", req, err)
	}
	if err != nil {
		return nil, err
	}
	return &types.MerchantUser{
		Id:          userInfo.Id,
		FirstName:   userInfo.FirstName,
		MidName:     userInfo.MidName,
		LastName:    userInfo.LastName,
		Password:    userInfo.Password,
		Telephone:   userInfo.Telephone,
		Mobliephone: userInfo.Mobliephone,
		MerchantId:  userInfo.MerchantId,
		Status:      userInfo.Status,
	}, nil
}
