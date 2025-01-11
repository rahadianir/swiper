package users

import (
	"context"

	"github.com/rahadianir/swiper/internal/models"
)

type UserRepositoryInterface interface {
	Register(ctx context.Context, user models.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetUserByUserID(ctx context.Context, id int) (models.User, error)
}

type UserLogicInterface interface {
	Register(ctx context.Context, payload models.RegisterRequest) (int, error)
	Login(ctx context.Context, payload models.LoginRequest) (string, string, error)
	GetProfileByID(ctx context.Context, id string) (models.User, error)
}

type AuthLogicInterface interface {
	GeneratePasswordHash(password string) (string, error)
	CompareHashAndPassword(hash string, password string) error
	GenerateJWT(payload models.User) (string, string, error)
}
