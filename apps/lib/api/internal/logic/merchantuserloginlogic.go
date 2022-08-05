package logic

import (
	"context"
	"strings"

	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/apps/lib/rpc/lib"
	"goecom/pkg/tool"
	"goecom/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdNullError = xerr.NewErrMsg("账号或密码不能为空")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")
var ErrUsernameNullError = xerr.NewErrMsg("账号不能为空")
var ErrEmailNullError = xerr.NewErrMsg("邮箱不能为空")
var ErrPasswordNullError = xerr.NewErrMsg("密码不能为空")
var ErrMerchantNameNullError = xerr.NewErrMsg("商户名不能为空")
var ErrMerchantAddError = xerr.NewErrMsg("商户加入错误")

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
func (l *MerchantuserloginLogic) Merchantuserlogin(req *types.MerchantUserLoginReq) (*types.MerchantUserLoginReply, error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.Wrapf(ErrUsernamePwdNullError, "Merchantuserlogin err null username:%s , password:%s, type:%s", req.UserName, req.Password, req.Type)
	}
	var resp types.MerchantUserLoginReply
	encodingPassword, errPass := tool.Md5ByString(req.Password)
	if errPass != nil {
		logx.Errorf("Merchantuserlogin encoding password error: %v", errPass)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin encoding password error :%s, err:%v", req.Password, errPass)
	}
	if req.Type == 1 {
		userUsernameEmail, errEmail := l.svcCtx.MerchantUserModel.FindOneByEmail(l.ctx, req.UserName)
		if errEmail != nil {
			logx.Errorf("merchant user login error: %v", errEmail)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.UserName, errEmail)
		}
		if encodingPassword != userUsernameEmail.Password {
			logx.Errorf("Merchantuserlogin password mismatch")
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin password mismatch username :%s, password:%s", req.UserName, req.Password)
		}
		tokenEmailResp, tokenEmailerr := l.svcCtx.LibRpc.GenerateToken(l.ctx, &lib.GenerateTokenReq{
			UserId: userUsernameEmail.Id,
		})
		if tokenEmailerr != nil {
			logx.Errorf("merchant user login error: %v", tokenEmailerr)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.UserName, tokenEmailerr)
		}
		resp.Id = userUsernameEmail.Id
		resp.NickName = userUsernameEmail.Nickname
		resp.UserName = userUsernameEmail.Username
		resp.Email = userUsernameEmail.Email
		resp.Password = userUsernameEmail.Password
		resp.Telephone = userUsernameEmail.Telephone
		resp.Mobliephone = userUsernameEmail.Mobliephone
		resp.MerchantId = userUsernameEmail.MerchantId
		resp.Status = userUsernameEmail.Status
		resp.CreatedAt = userUsernameEmail.CreatedAt
		resp.UpdatedAt = userUsernameEmail.UpdatedAt
		resp.AccessExpire = tokenEmailResp.AccessExpire
		resp.AccessToken = tokenEmailResp.AccessToken
		resp.RefreshAfter = tokenEmailResp.RefreshAfter

	} else {
		userUsernameUsername, errUsername := l.svcCtx.MerchantUserModel.FindOneByUsername(l.ctx, req.UserName)
		if errUsername != nil {
			logx.Errorf("merchant user login error: %v", errUsername)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.UserName, errUsername)
		}
		if encodingPassword != userUsernameUsername.Password {
			logx.Errorf("Merchantuserlogin password mismatch")
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin password mismatch username :%s, password:%s", req.UserName, req.Password)
		}
		tokenUsernameResp, tokenUsernameerr := l.svcCtx.LibRpc.GenerateToken(l.ctx, &lib.GenerateTokenReq{
			UserId: userUsernameUsername.Id,
		})
		if tokenUsernameerr != nil {
			logx.Errorf("merchant user login error: %v", tokenUsernameerr)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.UserName, tokenUsernameerr)
		}
		resp.Id = userUsernameUsername.Id
		resp.NickName = userUsernameUsername.Nickname
		resp.UserName = userUsernameUsername.Username
		resp.Email = userUsernameUsername.Email
		resp.Password = userUsernameUsername.Password
		resp.Telephone = userUsernameUsername.Telephone
		resp.Mobliephone = userUsernameUsername.Mobliephone
		resp.MerchantId = userUsernameUsername.MerchantId
		resp.Status = userUsernameUsername.Status
		resp.CreatedAt = userUsernameUsername.CreatedAt
		resp.UpdatedAt = userUsernameUsername.UpdatedAt
		resp.AccessExpire = tokenUsernameResp.AccessExpire
		resp.AccessToken = tokenUsernameResp.AccessToken
		resp.RefreshAfter = tokenUsernameResp.RefreshAfter
	}
	return &resp, nil
}
