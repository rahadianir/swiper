package premium

import "context"

type PremiumRepositoryInterface interface {
	EnablePremium(ctx context.Context, userID int) error
}

type PremiumLogicInterface interface {
	EnablePremium(ctx context.Context, userID int) error
}
