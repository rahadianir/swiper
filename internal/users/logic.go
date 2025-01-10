package users

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xerrors"
)

type UserLogic struct {
	*common.Dependencies
	UserRepo  UserRepositoryInterface
	AuthLogic AuthLogicInterface
}

func NewUserLogic(deps *common.Dependencies, userRepo UserRepositoryInterface, authLogic AuthLogicInterface) *UserLogic {
	return &UserLogic{
		Dependencies: deps,
		UserRepo:     userRepo,
		AuthLogic:    authLogic,
	}
}

func (logic *UserLogic) Register(ctx context.Context, payload models.RegisterRequest) (int, error) {

	// check for unwanted special characters
	// using this instead of regex because fuck regex lmao
	isContainSpecialChar := strings.ContainsFunc(payload.Password, func(r rune) bool {
		switch r {
		case '(', ')', '{', '}', '[', ']', '|', '`', '¬', '¦', '£', '$', '%', '^', '&', '*', '"', '<', '>', ':', ';', '#', '~', '_', '-', '+', '=', ',':
			return true
		default:
			return false
		}
	})
	if isContainSpecialChar {
		return 0, xerrors.ClientError{Err: fmt.Errorf("password cannot contains these special characters %s", "(){}[]|`¬¦\"£$%^&*<>:;#~_-+=,")}
	}

	pwHash, err := logic.AuthLogic.GeneratePasswordHash(payload.Password)
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate password hash", slog.Any("error", err))
	}

	user := models.User{
		Name:       payload.Name,
		Username:   payload.Username,
		Password:   pwHash,
		Age:        payload.Age,
		Location:   payload.Location,
		IsPremium:  false, // default
		IsVerified: false, // default
	}
	userID, err := logic.UserRepo.Register(ctx, user)
	if err != nil {
		// parse error
		if strings.Contains(err.Error(), "duplicate") {
			slog.WarnContext(ctx, "failed to register user", slog.Any("error", err))
			return 0, xerrors.ClientError{Err: fmt.Errorf("username is already used")}
		}

		// default
		slog.ErrorContext(ctx, "failed to register user", slog.Any("error", err))
		return 0, err
	}

	return userID, nil
}

func (logic *UserLogic) Login(ctx context.Context, payload models.LoginRequest) (string, string, error) {
	user, err := logic.UserRepo.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return "", "", err
	}

	err = logic.AuthLogic.CompareHashAndPassword(user.Password, payload.Password)
	if err != nil {
		return "", "", xerrors.ClientError{Err: fmt.Errorf("invalid password")}
	}

	token, refreshToken, err := logic.AuthLogic.GenerateJWT(user)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (logic *UserLogic) GetProfileByID(ctx context.Context, id string) (models.User, error) {
	user, err := logic.UserRepo.GetUserByUserID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
