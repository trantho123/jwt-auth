package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func (i impl) SighUp(signUpInput *model.SignUpInput) error {
	hashedPassword, err := utils.HashPassword(signUpInput.Password)
	if err != nil {
		return err
	}
	newUser := model.User{
		Email:            signUpInput.Email,
		Username:         signUpInput.UserName,
		Password:         hashedPassword,
		VerificationCode: utils.RandomString(30),
		Verified:         false,
		CreatedAt:        time.Now(),
	}

	// Is email exist in db
	isEmailExists, _ := i.repo.GetUserByEmail(context.Background(), newUser.Email)
	if isEmailExists.Email != "" {
		if !isEmailExists.Verified {
			return errors.New("please verify your email")
		} else {
			return errors.New("user with that email already exists")
		}
	} else {
		// Insert user
		insertErr := i.repo.InsertUser(context.Background(), &newUser)
		if insertErr != nil {
			return errors.New("failed to insert user")
		}
	}
	// send email
	user, getUserErr := i.repo.GetUserByEmail(context.Background(), newUser.Email)
	if getUserErr != nil {
		return errors.New("something bad happened")
	}
	id := user.ID.Hex()
	emailData := utils.EmailData{
		URL:      fmt.Sprintf("%s/signup/verify/%s/%s", os.Getenv("CLIENT_ORIGIN"), id, newUser.VerificationCode),
		UserName: newUser.Username,
		Subject:  "Your account verification code",
	}
	sendEmailErr := utils.SendEmail(&user, &emailData)
	if sendEmailErr != nil {
		return errors.New("failed to send email")
	}
	return nil
}

func (i impl) SignUpVerifyEmail(id, code string) error {
	user, getUserErr := i.repo.GetUserByID(context.Background(), id)
	if getUserErr != nil {
		return errors.New("something bad happened")
	}
	if !user.Verified {
		if user.VerificationCode == code {
			user.VerificationCode = ""
		}
	} else {
		return errors.New("this email has been confirmed")
	}
	user.Verified = true
	UpdateUserErr := i.repo.UpdateUser(context.Background(), &user)
	if UpdateUserErr != nil {
		return errors.New("something bad happened")
	}
	return nil
}
