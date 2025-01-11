package cache

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/redis/go-redis/v9"
)

type CacheStore struct {
	*common.Dependencies
}

func NewCacheStore(deps *common.Dependencies) *CacheStore {
	return &CacheStore{
		Dependencies: deps,
	}
}

func (c *CacheStore) Get(ctx context.Context, key string) (string, error) {
	res, err := c.RedisClient.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return res, err
	}
	return res, nil
}

func (c *CacheStore) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		valueBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = string(valueBytes)
	}
	res, err := c.RedisClient.Set(ctx, key, value, ttl).Result()
	log.Println(res, err)
	return err
}

func (c *CacheStore) Update(ctx context.Context, key string, value any) error {
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		valueBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		value = string(valueBytes)
	}
	res, err := c.RedisClient.SetArgs(ctx, key, value, redis.SetArgs{KeepTTL: true}).Result()
	log.Println(res, err)
	return err
}
