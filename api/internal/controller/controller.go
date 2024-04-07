package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/rds"
	"Jwtwithecdsa/api/internal/repository"
)

type Controller interface {
	SighUp(user *model.SignUpInput) error
	SignUpVerifyEmail(id, code string) error
	Login(loginInput *model.LoginInput) error
	LoginVerify(verifyOTP *model.VerifyOTP) (model.LoginResponse, error)
	RefreshAccessToken(refreshtoken string) (model.LoginResponse, error)
	GetMe(userid string) (model.UserResponse, error)
}
type impl struct {
	redis rds.RedisService
	repo  repository.Registry
}

func New(repo repository.Registry, rds rds.RedisService) Controller {
	return impl{
		repo:  repo,
		redis: rds}
}
