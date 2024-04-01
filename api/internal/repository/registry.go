package repository

import (
	"Jwtwithecdsa/api/internal/model"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Registry interface {
	InsertUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByID(ctx context.Context, id string) (model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type impl struct {
	mongoColl *mongo.Collection
}

func New(mongoColl *mongo.Collection) Registry {
	return impl{mongoColl: mongoColl}
}
