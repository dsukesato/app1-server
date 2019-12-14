package model

type Good struct {
	Id     int `json:"id"`
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

type Goods []Good

type PostGoodRequest struct {
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}
