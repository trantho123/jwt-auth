package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func createUserForTest(t *testing.T) *model.User {
	pass, err := utils.HashPassword(utils.RandomString(10))
	require.NoError(t, err)
	user := &model.User{
		Username:         utils.RandomString(5),
		Password:         pass,
		Email:            utils.RandomEmail(),
		Verified:         false,
		VerificationCode: "",
		CreatedAt:        time.Now(),
	}
	require.NoError(t, repo.InsertUser(context.Background(), user))

	return user
}
func TestInsertUser(t *testing.T) {
	createUserForTest(t)
}
func TestGetUser(t *testing.T) {
	user1 := createUserForTest(t)
	filter := bson.M{"username": user1.Username}
	user2, err := repo.GetUser(context.Background(), filter)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.VerificationCode, user2.VerificationCode)
	require.Equal(t, user1.Verified, user2.Verified)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createUserForTest(t)
	pass, err := utils.HashPassword(utils.RandomString(10))
	require.NoError(t, err)
	userForUpdate := &model.User{
		ID:               user1.ID,
		Username:         utils.RandomString(5),
		Password:         pass,
		Email:            utils.RandomEmail(),
		Verified:         false,
		VerificationCode: "",
		CreatedAt:        time.Now(),
	}
	require.NoError(t, repo.UpdateUser(context.Background(), userForUpdate))
}
