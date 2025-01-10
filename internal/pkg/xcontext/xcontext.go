package xcontext

import (
	"context"
	"fmt"
)

type contextKey string

const UserIDKey contextKey = "userID"

func GetUserID(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(UserIDKey).(int)
	if !ok {
		return 0, fmt.Errorf("user ID not found in context")
	}

	return userID, nil
}
