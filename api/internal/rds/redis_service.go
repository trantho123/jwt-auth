package rds

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService interface {
	SetOTP(ctx context.Context, email string, otp string, expiry time.Duration) error
	CompareOTP(ctx context.Context, email string, otp string) (bool, error)
}
type rds struct {
	client *redis.Client
}

func New(client *redis.Client) RedisService {
	return &rds{
		client: client,
	}
}

func (r *rds) SetOTP(ctx context.Context, email string, otp string, expiry time.Duration) error {
	err := r.client.Set(ctx, email, otp, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}
func (r *rds) CompareOTP(ctx context.Context, email string, otp string) (bool, error) {
	o, err := r.client.Get(ctx, email).Result()
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
