package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	InsertUser(ctx context.Context, user *model.User) error
}

type impl struct {
	mongoColl *mongo.Collection
}

func New(mongoColl *mongo.Collection) Registry {
	return impl{mongoColl: mongoColl}
}
