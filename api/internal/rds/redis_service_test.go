package rds

import (
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

var rdsTest RedisService

func TestMain(m *testing.M) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	rdsTest = New(rdb)
	os.Exit(m.Run())
}

func createRandomSetValue(t *testing.T) (string, string) {
	key := utils.RandomString(5)
	otp := utils.RandomString(10)
	err := rdsTest.Set(context.Background(), key, otp, time.Minute)
	require.NoError(t, err)
	return key, otp
}
func TestSet(t *testing.T) {
	createRandomSetValue(t)
}

func TestGet(t *testing.T) {
	key, otp := createRandomSetValue(t)
	val, err := rdsTest.Get(context.Background(), key)
	require.NoError(t, err)
	require.Equal(t, otp, val)
}

func TestCompare(t *testing.T) {

}

func TestDel(t *testing.T) {
	key, _ := createRandomSetValue(t)

	err := rdsTest.Del(context.Background(), key)
	require.NoError(t, err)
}
