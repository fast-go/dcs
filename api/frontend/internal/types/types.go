// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type IdentificationReq struct {
	Authorization string `json:"id"`
}

type UserinfoResp struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type SearchReq struct {
	Name string `form:"name"`
}

type SearchReply struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ProductGetDetailReq struct {
	ProductId int64 `json:"productId"`
}

type ProductGetDetailResp struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int64  `json:"stock"`
}

type CreateOrderReq struct {
	ProductId int64 `json:"productId"`
	Num       int64 `json:"num"`
}

type CreateOrderResp struct {
	Id          int64  `json:"id"`
	ProductName string `json:"productName"`
	ProductId   int64  `json:"productId"`
	Uid         int64  `json:"uid"`
	Num         int64  `json:"num"`
}
