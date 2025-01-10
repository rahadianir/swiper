package swiper

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xerrors"
)

type SwiperLogic struct {
	*common.Dependencies
	UserRepo   UserRepositoryInterface
	CacheStore CacheInterface
}

func NewSwiperLogic(deps *common.Dependencies, userRepo UserRepositoryInterface, cacheStore CacheInterface) *SwiperLogic {
	return &SwiperLogic{
		Dependencies: deps,
		UserRepo:     userRepo,
		CacheStore:   cacheStore,
	}
}

func (logic *SwiperLogic) GetTargetProfile(ctx context.Context, userID int) (models.User, error) {
	id := strconv.Itoa(userID)

	// get cached likes data
	cachedData, err := logic.CacheStore.Get(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	var cacheData models.LikesCache
	// only parse cache data if exists
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &cacheData)
		if err != nil {
			slog.ErrorContext(ctx, "failed to parse cache data", slog.Any("error", err))
			return models.User{}, err
		}
	}

	totalActivity := len(cacheData.Pass) + len(cacheData.Likes)
	if totalActivity >= 10 {
		return models.User{}, xerrors.ClientError{Err: fmt.Errorf("activity limit reached")}
	}

	viewedProfiles := append(cacheData.Pass, cacheData.Likes...)

	targetProfile, err := logic.UserRepo.GetRandomUser(ctx, viewedProfiles)
	if err != nil {
		return models.User{}, err
	}

	return targetProfile, nil
}

func (logic *SwiperLogic) SwipeRight(ctx context.Context, userID int, targetId int) error {
	return nil
}
