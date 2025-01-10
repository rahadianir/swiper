package swiper

import (
	"context"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
)

type SwiperLogic struct {
	*common.Dependencies
	UserRepo UserRepositoryInterface
}

func NewSwiperLogic(deps *common.Dependencies, userRepo UserRepositoryInterface) *SwiperLogic {
	return &SwiperLogic{
		Dependencies: deps,
		UserRepo:     userRepo,
	}
}

func (logic *SwiperLogic) GetTargetProfile(ctx context.Context, id int) (models.User, error) {

	return models.User{}, nil
}
