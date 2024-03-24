package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"context"
	"fmt"
	"log"
)

func (i impl) InsertUser(ctx context.Context, user *model.User) error {
	insertResult, err := i.mongoColl.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Inserted new user with ID:", insertResult.InsertedID)
	return nil
}
