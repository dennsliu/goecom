// Code generated by goctl. DO NOT EDIT.
package types

type Reply struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type Total struct {
	Total int64 `json:"total"`
}

type GetTokenReq struct {
}

type GetTokenReply struct {
	Code         int64  `json:"code"`
	Msg          string `json:"msg"`
	AccessToken  string `json:"accesstoken"`
	AccessExpire int64  `json:"accessexpire"`
	RefreshAfter int64  `json:"refreshafter"`
}

type Merchant struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}

type MerchantReply struct {
	Code      int64  `json:"code"`
	Msg       string `json:"msg"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}

type MerchantAddReq struct {
	Name   string `json:"name"`
	Status int64  `json:"status,optional,default=1"`
}

type MerchantUpdateReq struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status int64  `json:"status,optional,default=1"`
}

type MerchantGetReq struct {
	Id int64 `json:"id"`
}

type MerchantDeleteReq struct {
	Id int64 `json:"id"`
}

type MerchantDeleteReply struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Id   int64  `json:"id"`
}

type MerchantSearchReq struct {
	Keyword   string `json:"keyword,optional,default=''"`
	Page      int64  `json:"page,optional,default=1"`
	Status    int64  `json:"status,optional,default=1"`
	PageSize  int64  `json:"pagesize,optional,default=1"`
	LastId    int64  `json:"lastid,optional,default=0"`
	OrderType string `json:"ordertype,optional,default=desc"`
}

type MerchantSearchReply struct {
	Code        int64      `json:"code"`
	Msg         string     `json:"msg"`
	Mechants    []Merchant `json:"merchants"`
	IsEnd       bool       `json:"isend"`
	LastVal     int64      `json:"lastval"`
	Total       int64      `json:"total"`
	CurrentPage int64      `json:"currentpage"`
	TotalPage   int64      `json:"totalpage"`
}

type MerchantUser struct {
	Id          int64  `json:"id"`
	NickName    string `json:"nickname"`
	Email       string `json:"email"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Telephone   string `json:"telephone"`
	Mobliephone string `json:"mobliephone"`
	MerchantId  int64  `json:"merchantid"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
}

type MerchantUserReply struct {
	Code        int64  `json:"code"`
	Msg         string `json:"msg"`
	Id          int64  `json:"id"`
	NickName    string `json:"nickname"`
	Email       string `json:"email"`
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Telephone   string `json:"telephone"`
	Mobliephone string `json:"mobliephone"`
	MerchantId  int64  `json:"merchantid"`
	Status      int64  `json:"status"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
}

type MerchantUserAddReq struct {
	NickName    string `json:"nickname,optional,default=''"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Telephone   string `json:"telephone,optional,default=''"`
	Mobliephone string `json:"mobliephone,optional,default=''"`
	MerchantId  int64  `json:"merchantid"`
}

type MerchantUserUpdateReq struct {
	NickName    string `json:"nickname,optional,default=''"`
	Telephone   string `json:"telephone,optional,default=''"`
	Mobliephone string `json:"mobliephone,optional,default=''"`
	MerchantId  int64  `json:"merchantid"`
}

type MerchantUserGetReq struct {
	Id int64 `json:"id"`
}

type MerchantUserDeleteReq struct {
	Id int64 `json:"id"`
}

type MerchantUserDeleteReply struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Id   int64  `json:"id"`
}

type MerchantUserSearchReq struct {
	Keyword   string `json:"keyword,optional,default=''"`
	Status    int64  `json:"status,optional,default=1"`
	Page      int64  `json:"page,optional,default=1"`
	PageSize  int64  `json:"pagesize,optional,default=10"`
	LastId    int64  `json:"lastid,optional,default=0"`
	OrderType string `json:"ordertype,optional,default=desc"`
}

type MerchantUserSearchReply struct {
	Code         int64          `json:"code"`
	Msg          string         `json:"msg"`
	MechantUsers []MerchantUser `json:"merchantusers"`
	IsEnd        bool           `json:"isend"`
	LastVal      int64          `json:"lastval"`
	Total        int64          `json:"total"`
	CurrentPage  int64          `json:"currentpage"`
	TotalPage    int64          `json:"totalpage"`
}

type MerchantUserLoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Type     int64  `json:"type,optional,default=0"`
}

type MerchantUserLoginReply struct {
	Code         int64  `json:"code"`
	Msg          string `json:"msg"`
	Id           int64  `json:"id"`
	NickName     string `json:"nickname"`
	Email        string `json:"email"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Telephone    string `json:"telephone"`
	Mobliephone  string `json:"mobliephone"`
	MerchantId   int64  `json:"merchantid"`
	Status       int64  `json:"status"`
	CreatedAt    string `json:"createdat"`
	UpdatedAt    string `json:"updatedat"`
	AccessToken  string `json:"accesstoken"`
	AccessExpire int64  `json:"accessexpire"`
	RefreshAfter int64  `json:"refreshafter"`
}
