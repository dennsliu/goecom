package logic

import (
	"context"
	"errors"
	"strings"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"

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
		return nil, errors.New("Username and Password are required")
	}
	userInfo, err := l.svcCtx.MerchantUserModel.UserLogin(l.ctx, req.UserName)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("user not found username:" + req.UserName + ",password:" + req.Password)
	default:
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, errors.New("password mismatched")
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
		HttpCode:    200,
		Message:     "login successfully",
	}, nil
}
