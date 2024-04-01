package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	InsertUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, filter bson.M) (model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type impl struct {
	mongoColl *mongo.Collection
}

func New(mongoColl *mongo.Collection) Registry {
	return impl{mongoColl: mongoColl}
}
