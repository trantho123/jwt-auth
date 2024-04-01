package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (i impl) InsertUser(ctx context.Context, user *model.User) error {
	_, err := i.mongoColl.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (i impl) UpdateUser(ctx context.Context, user *model.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"email":            user.Email,
		"username":         user.Username,
		"password":         user.Password,
		"verificationcode": user.VerificationCode,
		"verified":         user.Verified,
		"createdat":        user.CreatedAt,
	}}
	_, err := i.mongoColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user in MongoDB: %v", err)
	}

	return nil
}

func (i impl) GetUser(ctx context.Context, filter bson.M) (model.User, error) {
	var user model.User
	err := i.mongoColl.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, err
		}
		return model.User{}, err
	}
	return user, nil
}
