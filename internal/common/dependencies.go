package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/rahadianir/swiper/internal/config"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	Config      *config.Config
	DB          *sqlx.DB
	RedisClient *redis.Client
	Validator   *validator.Validate
}
