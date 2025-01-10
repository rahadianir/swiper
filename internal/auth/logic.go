package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthLogic struct {
	*common.Dependencies
}

func NewAuthLogic(deps *common.Dependencies) *AuthLogic {
	return &AuthLogic{
		Dependencies: deps,
	}
}

func (logic *AuthLogic) GeneratePasswordHash(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashBytes), nil
}

func (logic *AuthLogic) CompareHashAndPassword(hash string, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (logic *AuthLogic) GenerateJWT(payload models.User) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Issuer:    "swiper",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		UserID: payload.ID,
	})

	tokenString, err := token.SignedString([]byte(logic.Config.JWTSecret))
	if err != nil {
		return "", "", err
	}

	// TODO: add refresh token
	return tokenString, "", nil
}
