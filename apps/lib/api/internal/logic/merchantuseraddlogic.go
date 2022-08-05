package logic

import (
	"context"
	"strings"
	"time"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/model"
	"goecom/pkg/tool"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantuseraddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantuseraddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantuseraddLogic {
	return &MerchantuseraddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *MerchantuseraddLogic) Merchantuseradd(req *types.MerchantUserAddReq) (resp *types.MerchantUser, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.UserName)) == 0 {
		logx.Errorf("Merchantuseradd err null username:%s", req.UserName)
		return nil, errors.Wrapf(ErrUsernameNullError, "Merchantuseradd err null username:%s", req.UserName)
	}
	if len(strings.TrimSpace(req.Email)) == 0 {
		logx.Errorf("Merchantuseradd err null email:%s", req.Email)
		return nil, errors.Wrapf(ErrEmailNullError, "Merchantuseradd err null email:%s", req.Email)
	}
	if len(strings.TrimSpace(req.Password)) == 0 {
		logx.Errorf("Merchantuseradd err null password:%s", req.Password)
		return nil, errors.Wrapf(ErrPasswordNullError, "Merchantuseradd err null password:%s", req.Password)
	}
	encodingPassword, errPass := tool.Md5ByString(req.Password)
	if errPass != nil {
		logx.Errorf("Merchantuseradd encoding password error: %v", errPass)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuseradd encoding password error :%s, err:%v", req.Password, errPass)
	}
	merchantUserModel := new(model.MerchantUser)
	merchantUserModel.Nickname = req.NickName
	merchantUserModel.Username = req.UserName
	merchantUserModel.Email = req.Email
	merchantUserModel.Password = encodingPassword
	merchantUserModel.Telephone = req.Telephone
	merchantUserModel.Mobliephone = req.Mobliephone
	merchantUserModel.MerchantId = req.MerchantId
	merchantUserModel.Status = 1
	merchantUserModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	merchantUserModel.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err = l.svcCtx.MerchantUserModel.Insert(l.ctx, merchantUserModel)
	if err != nil {
		logx.Errorf("Merchantuseradd insert error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuseradd insert err username:%s, err:%v", req.UserName, err)
	}
	userInfo, err := l.svcCtx.MerchantUserModel.FindOneByUsername(l.ctx, req.UserName)
	if err != nil {
		logx.Errorf("Merchantuseradd find error: %v", err)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuseradd find err username:%s, err:%v", req.UserName, err)
	}
	return &types.MerchantUser{
		Id:          userInfo.Id,
		NickName:    userInfo.Nickname,
		UserName:    userInfo.Username,
		Email:       userInfo.Email,
		Password:    userInfo.Password,
		Telephone:   userInfo.Telephone,
		Mobliephone: userInfo.Mobliephone,
		MerchantId:  userInfo.MerchantId,
		Status:      userInfo.Status,
		CreatedAt:   userInfo.CreatedAt,
		UpdatedAt:   userInfo.UpdatedAt,
	}, nil
}
