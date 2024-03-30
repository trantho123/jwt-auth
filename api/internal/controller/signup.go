package controller

import (
	"Jwtwithecdsa/api/internal/model"
	"Jwtwithecdsa/api/internal/utils"
	"context"
	"errors"
	"time"
)

func (i impl) SighUp(signUpInput *model.SignUpInput) error {
	hashedPassword, err := utils.HashPassword(signUpInput.Password)
	if err != nil {
		return err
	}
	newUser := model.User{
		Email:     signUpInput.Email,
		Username:  signUpInput.UserName,
		Password:  hashedPassword,
		Verified:  false,
		CreatedAt: time.Now(),
	}
	// Generate otp
	newOTP, newOTPErr := utils.GetRandNum()
	if newOTPErr != nil {
		return errors.New("failed to gen otp email: " + newOTPErr.Error())
	}
	// Is email exist in db
	isEmailExists, _ := i.repo.GetUser(context.Background(), newUser.Email)
	if isEmailExists.Email != "" {
		if !isEmailExists.Verified {
			sendEmailErr := utils.SendEmail(newUser.Email, "Confirm OTP", newOTP)
			if sendEmailErr != nil {
				return errors.New("failed to send email: " + sendEmailErr.Error())
			}
		} else {
			return errors.New("user with that email already exists")
		}
	} else {
		// Insert user
		insertErr := i.repo.InsertUser(context.Background(), &newUser)
		if insertErr != nil {
			return errors.New("failed to insert user: " + insertErr.Error())
		}
	}
	// save in redis
	cacheErr := i.rds.SetOTP(context.Background(), newUser.Email, newOTP, 5*time.Minute)
	if cacheErr != nil {
		return cacheErr
	}
	// send email
	sendEmailErr := utils.SendEmail(newUser.Email, "Confirm OTP", newOTP)
	if sendEmailErr != nil {
		return errors.New("failed to send email: " + sendEmailErr.Error())
	}
	return nil
}

func (i impl) SignUpVerifyEmail(verify *model.VerifyEmail) error {
	// Compare OTP
	isSameOtp, err := i.rds.CompareOTP(context.Background(), verify.Email, verify.OTPCode)
	if err != nil {
		return err
	}
	if !isSameOtp {
		return errors.New("invalid verification code")
	}
	// Get User in mongodb
	user, getUserErr := i.repo.GetUser(context.Background(), verify.Email)
	if getUserErr != nil {
		return errors.New("something bad happened")
	}
	// Update veridied
	if user.Verified {
		return errors.New("this email has been confirmed")
	}
	user.Verified = true
	UpdateUserErr := i.repo.UpdateUser(context.Background(), &user)
	if UpdateUserErr != nil {
		return err
	}
	return nil
}
