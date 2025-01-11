package swiper

import (
	"context"
	"time"

	"github.com/rahadianir/swiper/internal/models"
)
//go:generate mockgen -source port.go -destination mock/mock_port.go
type SwiperLogicInterface interface {
	GetTargetProfile(ctx context.Context, userID int) (models.User, error)
	SwipeRight(ctx context.Context, userID int, targetId int) (bool, error)
	SwipeLeft(ctx context.Context, userID int, targetId int) error
}

type SwiperRepositoryInterface interface {
	StoreUserLike(ctx context.Context, userID int, targetID int) error
	GetUserLikedUserIDs(ctx context.Context, userID int, params models.LikedUserParams) ([]int, error)
	UpdateMatchStatus(ctx context.Context, userID int, targetID int, status bool) error
}

type UserRepositoryInterface interface {
	GetUserByUserID(ctx context.Context, id int) (models.User, error)
	GetRandomUser(ctx context.Context, excludeList []int) (models.User, error)
}

type CacheInterface interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Update(ctx context.Context, key string, value any) error
}
