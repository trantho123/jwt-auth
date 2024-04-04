package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"errors"
	"time"
)

func (i impl) RefreshAccessToken(refreshtoken string) (model.LoginResponse, error) {
	_, err := utils.VerifyToken(refreshtoken)
	if err != nil {
		return model.LoginResponse{}, err
	}
	userid, err := i.redis.Get(context.Background(), refreshtoken)
	if err != nil {
		return model.LoginResponse{}, err
	}
	// Create new accesstoken and refreshtoken
	newAccessTokenStr, err := utils.GenerateToken(userid, 15)
	if err != nil {
		return model.LoginResponse{}, errors.New("cant refresh, please login again")
	}
	newRefreshTokenStr, err := utils.GenerateToken(userid, 60)
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
