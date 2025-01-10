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

	var cacheData models.ActivityCache
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

	viewedProfiles := []int{userID} // to prevent viewing user's own profile
	viewedProfiles = append(viewedProfiles, cacheData.Pass...)
	viewedProfiles = append(viewedProfiles, cacheData.Likes...)

	targetProfile, err := logic.UserRepo.GetRandomUser(ctx, viewedProfiles)
	if err != nil {
		return models.User{}, err
	}

	return targetProfile, nil
}

func (logic *SwiperLogic) SwipeRight(ctx context.Context, userID int, targetId int) error {
	// id := strconv.Itoa(userID)
	targetID := strconv.Itoa(targetId)

	// check match
	isMatch := checkMatch(targetId, []int{})
	if isMatch {
		// TODO: broadcast match and initiate chat room or some shit
		slog.InfoContext(ctx, "ciee match", slog.Any("user_id", userID), slog.Any("target_id", targetID))
	}

	return nil
}

func checkMatch(userID int, likedBy []int) bool {
	for _, id := range likedBy {
		if id == userID {
			return true
		}
	}

	return false
}
