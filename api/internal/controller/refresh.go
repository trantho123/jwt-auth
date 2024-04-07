package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"errors"
	"os"
	"time"
)

func (i impl) RefreshAccessToken(refreshtoken string) (model.LoginResponse, error) {
	_, err := utils.VerifyToken(refreshtoken, os.Getenv("PUBLIC_REFRESH_KEY"))
	if err != nil {
		return model.LoginResponse{}, err
	}
	userid, err := i.redis.Get(context.Background(), refreshtoken)
	if err != nil {
		return model.LoginResponse{}, err
	}
	// Create new accesstoken and refreshtoken
	newAccessTokenStr, err := utils.GenerateToken(userid, os.Getenv("PRIVATE_ACCESS_KEY"), 15)
	if err != nil {
		return model.LoginResponse{}, errors.New("cant refresh, please login again")
	}
	newRefreshTokenStr, err := utils.GenerateToken(userid, os.Getenv("PRIVATE_REFRESH_KEY"), 60)
	if err != nil {
		return model.LoginResponse{}, errors.New("cant refresh, please login again")
	}
	if err := i.redis.Set(context.Background(), newRefreshTokenStr, userid, 60*time.Minute); err != nil {
		return model.LoginResponse{}, errors.New("cant refresh, please login again")
	}
	res := model.LoginResponse{
		AccessToken:  newAccessTokenStr,
		RefreshToken: newRefreshTokenStr,
	}
	return res, nil
}
