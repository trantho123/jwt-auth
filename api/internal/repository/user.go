package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (i impl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := i.mongoColl.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}
	return user, nil
}

func (i impl) UpdateUser(ctx context.Context, user *model.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"verified": user.Verified,
	}}
	_, err := i.mongoColl.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user in MongoDB: %v", err)
	}

	return nil
}

func (i impl) GetUserByID(ctx context.Context, id string) (model.User, error) {
	userID, _ := primitive.ObjectIDFromHex(id)
	var user model.User
	err := i.mongoColl.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}
	return user, nil
}
