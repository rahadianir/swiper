package swiper

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xerrors"
)

type SwiperLogic struct {
	*common.Dependencies
	UserRepo   UserRepositoryInterface
	CacheStore CacheInterface
	SwiperRepo SwiperRepositoryInterface
}

func NewSwiperLogic(deps *common.Dependencies, userRepo UserRepositoryInterface, cacheStore CacheInterface, swiperRepo SwiperRepositoryInterface) *SwiperLogic {
	return &SwiperLogic{
		Dependencies: deps,
		UserRepo:     userRepo,
		CacheStore:   cacheStore,
		SwiperRepo:   swiperRepo,
	}
}

func (logic *SwiperLogic) GetTargetProfile(ctx context.Context, userID int) (models.User, error) {
	// convert id from int to string for cachestore
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

	// get matched profile so it won't be shown anymore
	matchedIDs, err := logic.SwiperRepo.GetUserLikedUserIDs(ctx, userID, models.LikedUserParams{
		IsMatched: true,
	})
	if err != nil {
		return models.User{}, fmt.Errorf("failed to get matched ids to exclude")
	}

	excludedIDs := []int{userID} // to prevent viewing user's own profile
	excludedIDs = append(excludedIDs, cacheData.Pass...)
	excludedIDs = append(excludedIDs, cacheData.Likes...)
	excludedIDs = append(excludedIDs, matchedIDs...)

	// get random user with excluded list
	targetProfile, err := logic.UserRepo.GetRandomUser(ctx, excludedIDs)
	if err != nil {
		return models.User{}, err
	}

	return targetProfile, nil
}

func (logic *SwiperLogic) SwipeRight(ctx context.Context, userID int, targetId int) (bool, error) {
	// convert id from int to string for cachestore
	id := strconv.Itoa(userID)
	targetID := strconv.Itoa(targetId)

	// get cached likes data
	cachedData, err := logic.CacheStore.Get(ctx, id)
	if err != nil {
		return false, err
	}

	var cacheData models.ActivityCache
	isUserNewActivity := true
	// only parse cache data if exists
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &cacheData)
		if err != nil {
			slog.ErrorContext(ctx, "failed to parse cache data", slog.Any("error", err))
			return false, err
		}
		isUserNewActivity = false
	}

	// if user already swiped 10x in 24 hours, return error
	// TODO: bypass this with premium
	totalActivity := len(cacheData.Pass) + len(cacheData.Likes)
	if totalActivity >= 10 {
		return false, xerrors.ClientError{Err: fmt.Errorf("activity limit reached")}
	}

	// store likes both in database and cache
	err = logic.SwiperRepo.StoreUserLike(ctx, userID, targetId)
	if err != nil {
		return false, err
	}

	cacheData.Likes = append(cacheData.Likes, targetId)
	if isUserNewActivity {
		err := logic.CacheStore.Set(ctx, id, cacheData, (24 * time.Hour))
		if err != nil {
			slog.WarnContext(ctx, "failed to store cache data", slog.Any("error", err))
		}
	} else {
		err := logic.CacheStore.Update(ctx, id, cacheData)
		if err != nil {
			slog.WarnContext(ctx, "failed to updates cache data", slog.Any("error", err))
		}
	}

	// check match
	// get target likes
	targetLikes, err := logic.SwiperRepo.GetUserLikedUserIDs(ctx, targetId, models.LikedUserParams{})
	if err != nil {
		return false, err
	}

	// check whether userid is in target likes
	isMatch := checkMatch(userID, targetLikes)
	if isMatch {
		// TODO: broadcast match and initiate chat room or some shit
		slog.InfoContext(ctx, "ciee match", slog.Any("user_id", userID), slog.Any("target_id", targetID))
	}

	return isMatch, nil
}

func checkMatch(userID int, likedBy []int) bool {
	for _, id := range likedBy {
		if id == userID {
			return true
		}
	}

	return false
}

func (logic *SwiperLogic) SwipeLeft(ctx context.Context, userID int, targetId int) error {
	// convert id from int to string for cachestore
	id := strconv.Itoa(userID)

	// get cached likes data
	cachedData, err := logic.CacheStore.Get(ctx, id)
	if err != nil {
		return err
	}

	var cacheData models.ActivityCache
	isUserNewActivity := true
	// only parse cache data if exists
	if cachedData != "" {
		err = json.Unmarshal([]byte(cachedData), &cacheData)
		if err != nil {
			slog.ErrorContext(ctx, "failed to parse cache data", slog.Any("error", err))
			return err
		}
		isUserNewActivity = false
	}

	// if user already swiped 10x in 24 hours, return error
	// TODO: bypass this with premium
	totalActivity := len(cacheData.Pass) + len(cacheData.Likes)
	if totalActivity >= 10 {
		return xerrors.ClientError{Err: fmt.Errorf("activity limit reached")}
	}

	cacheData.Pass = append(cacheData.Pass, targetId)
	if isUserNewActivity {
		err := logic.CacheStore.Set(ctx, id, cacheData, (24 * time.Hour))
		if err != nil {
			slog.WarnContext(ctx, "failed to store cache data", slog.Any("error", err))
		}
	} else {
		err := logic.CacheStore.Update(ctx, id, cacheData)
		if err != nil {
			slog.WarnContext(ctx, "failed to updates cache data", slog.Any("error", err))
		}
	}

	return nil
}
