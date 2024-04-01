package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/rds"
	"Jwtwithecdsa/api/internal/repository"
)

type Controller interface {
	SighUp(user *model.SignUpInput) error
	SignUpVerifyEmail(id, code string) error
}
type impl struct {
	rds  rds.RedisService
	repo repository.Registry
}

func New(repo repository.Registry, rds rds.RedisService) Controller {
	return impl{
		repo: repo,
		rds:  rds}
}
