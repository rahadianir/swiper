package premium

import (
	"context"

	"github.com/rahadianir/swiper/internal/common"
)

type PremiumRepo struct {
	*common.Dependencies
}

func NewPremiumRepo(deps *common.Dependencies) *PremiumRepo {
	return &PremiumRepo{
		Dependencies: deps,
	}
}

func (r *PremiumRepo) EnablePremium(ctx context.Context, userID int) error {
	q := `UPDATE users
			SET 
				is_premium = true,
				is_verified = true,
				updated_at = now()
			WHERE
				users.id = $1`
	_, err := r.DB.ExecContext(ctx, q, userID)
	if err != nil {
		return err
	}
	return nil
}
