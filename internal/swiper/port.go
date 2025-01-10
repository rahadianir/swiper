package swiper

import (
	"context"

	"github.com/rahadianir/swiper/internal/models"
)

type UserRepositoryInterface interface {
	GetUserByUserID(ctx context.Context, id int) (models.User, error)
}
