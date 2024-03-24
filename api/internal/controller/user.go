package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"context"
)

func (i impl) SighUp(user *model.User) error {
	err := i.repo.InsertUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
