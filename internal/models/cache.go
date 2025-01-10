package models

type ActivityCache struct {
	UserID int   `json:"id"`
	Pass   []int `json:"pass"`
	Likes  []int `json:"likes"`
}
