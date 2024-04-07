package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (i impl) Login(loginInput *model.LoginInput) error {
	// Is user exist
	filter := bson.M{"username": loginInput.UserName}
	user, err := i.repo.GetUser(context.Background(), filter)
	if err != nil {
		return errors.New("user not found")
	}
	if !user.Verified {
		return errors.New("please verify your email address before logging in")
	}
	// verify pass
	if err := utils.VerifyPassword(user.Password, loginInput.Password); err != nil {
		return errors.New("wrong password")
	}
	if !user.Verified {
		return errors.New("account not verified")
	}
	// random otp
	key := user.ID.Hex()
	otp, _ := utils.GetRandNum()
	setOTPErr := i.redis.Set(context.Background(), key, otp, 5*time.Minute)
	if setOTPErr != nil {
		return errors.New("something bad happened")
	}
	// send email
	emailData := utils.EmailData{
		Subject: "Your OTP for login",
		Content: fmt.Sprintf("Your OTP is:%s", otp),
	}
	sendEmailErr := utils.SendEmail(&user, &emailData)
	if sendEmailErr != nil {
		return errors.New("failed to send email")
	}
	return nil
}

func (i impl) LoginVerify(verifyOTP *model.VerifyOTP) (model.LoginResponse, error) {
	// is email exist
	filter := bson.M{"email": verifyOTP.Email}
	user, err := i.repo.GetUser(context.Background(), filter)
	if err != nil {
		return model.LoginResponse{}, errors.New("user not found")
	}
	// compare otp in redis
	isOTPSame, compareOtpErr := i.redis.Compare(context.Background(), user.ID.Hex(), verifyOTP.OTPCode)
	if compareOtpErr != nil {
		return model.LoginResponse{}, errors.New("something bad happened")
	}
	if !isOTPSame {
		return model.LoginResponse{}, errors.New("invalid OTP")
	}
	// create token
	accessTokenStr, err := utils.GenerateToken(user.ID.Hex(), os.Getenv("PRIVATE_ACCESS_KEY"), 15)
	if err != nil {
		return model.LoginResponse{}, errors.New("something bad happened")
	}
	refreshTokenStr, err := utils.GenerateToken(user.ID.Hex(), os.Getenv("PRIVATE_REFRESH_KEY"), 60)
	if err != nil {
		return model.LoginResponse{}, errors.New("something bad happened")
	}
	if err := i.redis.Set(context.Background(), refreshTokenStr, user.ID.Hex(), 60*time.Minute); err != nil {
		fmt.Sprintln("set otp :", err)
		return model.LoginResponse{}, errors.New("something bad happened")
	}
	res := model.LoginResponse{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}
	return res, nil
}
