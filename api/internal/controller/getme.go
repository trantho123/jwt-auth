package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (i impl) GetMe(userid string) (model.UserResponse, error) {
	// Is user exist
	objId, _ := primitive.ObjectIDFromHex(userid)
	user, err := i.repo.GetUser(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return model.UserResponse{}, errors.New("something bad happened")
	}
	useRes := model.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
	return useRes, nil
}
