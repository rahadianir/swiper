package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	"github.com/rahadianir/swiper/internal/pkg/xcontext"
	"github.com/rahadianir/swiper/internal/pkg/xhttp"
)

type AuthMiddleware struct {
	*common.Dependencies
}

func (m *AuthMiddleware) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			xhttp.SendJSONResponse(w, models.BaseResponse{
				Error:   "empty authorization bearer token",
				Message: "unauthorized!",
			}, http.StatusUnauthorized)
			return
		}

		// parse bearer token
		tempString := strings.Split(bearerToken, " ")
		if len(tempString) < 2 {
			xhttp.SendJSONResponse(w, models.BaseResponse{
				Error:   "invalid authorization bearer token format",
				Message: "unauthorized!",
			}, http.StatusUnauthorized)
			return
		}

		tokenString := tempString[1]
		token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(m.Config.JWTSecret), nil
		})
		if err != nil {
			xhttp.SendJSONResponse(w, models.BaseResponse{
				Error:   err.Error(),
				Message: "invalid token",
			}, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*models.CustomClaims)
		if !ok {
			xhttp.SendJSONResponse(w, models.BaseResponse{
				Error:   "invalid token claims",
				Message: "invalid token",
			}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), xcontext.UserIDKey, claims.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
