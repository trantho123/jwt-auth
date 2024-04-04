package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (i impl) Login(loginInput *model.LoginInput) error {
	filter := bson.M{"username": loginInput.UserName}
	user, err := i.repo.GetUser(context.Background(), filter)
	if err != nil {
		return errors.New("user not found")
	}
	if err := utils.VerifyPassword(user.Password, loginInput.Password); err != nil {
		return errors.New("wrong password")
	}
	if !user.Verified {
		return errors.New("account not verified")
	}
	key := user.ID.Hex()
	otp, _ := utils.GetRandNum()
	setOTPErr := i.redis.SetOTP(context.Background(), key, otp, 5*time.Minute)
	if setOTPErr != nil {
		return errors.New("something bad happened")
	}
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
