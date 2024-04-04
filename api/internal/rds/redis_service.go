package rds

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	Set(ctx context.Context, key string, otp string, expiry time.Duration) error
	Compare(ctx context.Context, key string, otp string) (bool, error)
	Del(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (string, error)
}
type rds struct {
	client *redis.Client
}

func New(client *redis.Client) RedisService {
	return &rds{
		client: client,
	}
}

func (r *rds) Set(ctx context.Context, key string, otp string, expiry time.Duration) error {
	err := r.client.Set(ctx, key, otp, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *rds) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil || err != nil {
		return "", errors.New("Cant get userid in redis")
	}
	return result, nil
}

func (r *rds) Compare(ctx context.Context, key string, otp string) (bool, error) {
	o, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, err
	} else if err != nil {
		return false, err
	}
	if o != otp {
		return false, nil
	}
	return true, nil
}

func (r *rds) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
