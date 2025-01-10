package swiper

import (
	"context"

	"github.com/rahadianir/swiper/internal/common"
)

type SwipeRepo struct {
	*common.Dependencies
}

func NewSwipeRepo(deps *common.Dependencies) *SwipeRepo {
	return &SwipeRepo{
		Dependencies: deps,
	}
}

func (r *SwipeRepo) GetUserLikedUserIDs(ctx context.Context, userID int) ([]int, error) {

	return []int{}, nil
}
