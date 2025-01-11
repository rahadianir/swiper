package models

import "time"

type User struct {
	ID         int        `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Username   string     `json:"username,omitempty"`
	Password   string     `json:"password,omitempty"`
	Age        int        `json:"age,omitempty"`
	Gender     string     `json:"gender,omitempty"`
	Location   string     `json:"string,omitempty"`
	IsPremium  bool       `json:"is_premium,omitempty"`
	IsVerified bool       `json:"is_verified,omitempty"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,alpha"`
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
	Age      int    `json:"age" validate:"required,number,gte=0,lte=100"`
	Gender   string `json:"gender" validate:"oneof=male female"`
	Location string `json:"location" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
