package models

type LikesCache struct {
	UserID  int   `json:"id"`
	Pass    []int `json:"pass"`
	Likes   []int `json:"likes"`
	LikedBy []int `json:"liked_by"`
}
