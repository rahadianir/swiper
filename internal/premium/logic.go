package premium

import (
	"context"

	"github.com/rahadianir/swiper/internal/common"
)

type PremiumLogic struct {
	*common.Dependencies
	PremiumRepo PremiumRepositoryInterface
}

func NewPremiumLogic(deps *common.Dependencies, premiumRepo PremiumRepositoryInterface) *PremiumLogic {
	return &PremiumLogic{
		Dependencies: deps,
		PremiumRepo:  premiumRepo,
	}
}

func (logic *PremiumLogic) EnablePremium(ctx context.Context, userID int) error {
	return logic.PremiumRepo.EnablePremium(ctx, userID)
}
