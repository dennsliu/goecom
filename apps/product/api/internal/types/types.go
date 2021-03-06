// Code generated by goctl. DO NOT EDIT.
package types

type Brand struct {
	Id         int64  `json:"id"`
	Order      int64  `json:"order"`
	Status     int64  `json:"status"`
	MerchantId int64  `json:"merchantid"`
	ParentId   int64  `json:"parentid"`
	Code       string `json:"code"`
	CreatedAt  string `json:"createdat"`
	UpdateAt   string `json:"updateat"`
	HttpCode   int64  `json:"httpcode"`
	Message    string `json:"message"`
}

type BrandAddReq struct {
	Order      int64  `json:"order"`
	Status     int64  `json:"status"`
	MerchantId int64  `json:"merchantid"`
	ParentId   int64  `json:"parentid"`
	Code       string `json:"code"`
}

type BrandUpdateReq struct {
	Id         int64  `json:"id"`
	Order      int64  `json:"order"`
	Status     int64  `json:"status"`
	MerchantId int64  `json:"merchantid"`
	ParentId   int64  `json:"parentid"`
	Code       string `json:"code"`
}

type BrandGetReq struct {
	Id int64 `json:"id"`
}

type BrandDeleteReq struct {
	Id string `json:"id"`
}

type BrandDeleteReply struct {
	HttpCode int64  `json:"httpcode"`
	Message  string `json:"message"`
}

type BrandSearchReq struct {
	Keyword  string `json:"keyword"`
	Page     int64  `json:"page"`
	Status   int64  `json:"page"`
	PageSize int64  `json:"pagesize"`
	OrderBy  string `json:"orderby"`
}

type BrandSearchReply struct {
	Brands   []Brand `json:"brands"`
	IsEnd    bool    `json:"isend"`
	LastVal  int64   `json:"lastval"`
	HttpCode int64   `json:"httpcode"`
	Message  string  `json:"message"`
}
