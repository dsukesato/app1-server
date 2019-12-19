package model

type Good struct {
	Id     int  `json:"id"`
	PostId int  `json:"post_id"`
	UserId int  `json:"user_id"`
	State  bool `json:"state"`
}

type Goods []Good

type PostGoodRequest struct {
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

type PostGoodResponse struct {
	Id       int  `json:"id"`
	PostId   int  `json:"post_id"`
	UserId   int  `json:"user_id"`
	State    bool `json:"state"`
	PostGood int  `json:"post_good"`
}
