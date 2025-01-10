package swiper

import (
	"context"
	"time"

	"github.com/rahadianir/swiper/internal/models"
)

type SwiperLogicInterface interface {
	GetTargetProfile(ctx context.Context, userID int) (models.User, error)
}

type UserRepositoryInterface interface {
	GetUserByUserID(ctx context.Context, id string) (models.User, error)
	GetRandomUser(ctx context.Context, excludeList []int) (models.User, error)
}

type CacheInterface interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
}
