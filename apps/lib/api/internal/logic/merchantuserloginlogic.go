package logic

import (
	"context"
	"fmt"
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
	fmt.Print("--------------777-------------")
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.Wrapf(ErrUsernamePwdNullError, "Merchantuserlogin err null username:%s , password:%s, type:%s", req.Username, req.Password, req.Type)
	}
	var resp types.MerchantUserLoginReply
	encodingPassword, errPass := tool.Md5ByString(req.Password)
	if errPass != nil {
		logx.Errorf("Merchantuserlogin encoding password error: %v", errPass)
		return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin encoding password error :%s, err:%v", req.Password, errPass)
	}
	if req.Type == 1 {
		userUsernameEmail, errEmail := l.svcCtx.MerchantUserModel.FindOneByEmail(l.ctx, req.Username)
		if errEmail != nil {
			logx.Errorf("merchant user login error: %v", errEmail)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.Username, errEmail)
		}
		if encodingPassword != userUsernameEmail.Password {
			logx.Errorf("Merchantuserlogin password mismatch")
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin password mismatch username :%s, password:%s", req.Username, req.Password)
		}
		tokenEmailResp, tokenEmailerr := l.svcCtx.LibRpc.GenerateToken(l.ctx, &lib.GenerateTokenReq{
			UserId: userUsernameEmail.Id,
		})
		if tokenEmailerr != nil {
			logx.Errorf("merchant user login error: %v", tokenEmailerr)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.Username, tokenEmailerr)
		}
		resp.Code = 200
		resp.Msg = "login successfully"
		resp.Id = userUsernameEmail.Id
		resp.Nickname = userUsernameEmail.Nickname
		resp.Username = userUsernameEmail.Username
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
		userUsernameUsername, errUsername := l.svcCtx.MerchantUserModel.FindOneByUsername(l.ctx, req.Username)
		if errUsername != nil {
			logx.Errorf("merchant user login error: %v", errUsername)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.Username, errUsername)
		}
		if encodingPassword != userUsernameUsername.Password {
			logx.Errorf("Merchantuserlogin password mismatch")
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin password mismatch username :%s, password:%s", req.Username, req.Password)
		}
		tokenUsernameResp, tokenUsernameerr := l.svcCtx.LibRpc.GenerateToken(l.ctx, &lib.GenerateTokenReq{
			UserId: userUsernameUsername.Id,
		})
		if tokenUsernameerr != nil {
			logx.Errorf("merchant user login error: %v", tokenUsernameerr)
			return nil, errors.Wrapf(ErrUsernamePwdError, "Merchantuserlogin err :%s, err:%v", req.Username, tokenUsernameerr)
		}
		fmt.Print("--------------99999-------------")
		resp.Code = 200
		resp.Msg = "login successfully"
		resp.Id = userUsernameUsername.Id
		resp.Nickname = userUsernameUsername.Nickname
		resp.Username = userUsernameUsername.Username
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
