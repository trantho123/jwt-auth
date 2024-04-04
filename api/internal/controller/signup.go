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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Is username exist in db
	isUserNameExist, _ := i.repo.GetUser(context.Background(), bson.M{"username": newUser.Username})
	if isUserNameExist.Username != "" {
		return fmt.Errorf("username '%s' already exists", newUser.Username)
	}
	// Is email exist in db
	isEmailExists, _ := i.repo.GetUser(context.Background(), bson.M{"email": newUser.Email})
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
	user, getUserErr := i.repo.GetUser(context.Background(), bson.M{"email": newUser.Email})
	if getUserErr != nil {
		return errors.New("something bad happened")
	}
	id := user.ID.Hex()
	url := fmt.Sprintf("%s/signup/verify/%s/%s", os.Getenv("CLIENT_ORIGIN"), id, newUser.VerificationCode)
	emailData := utils.EmailData{
		URL:     url,
		Subject: "Your account verification code",
		Content: fmt.Sprintf(`Hello %s,<br/>
		Thank you for registering with us!<br/>
		Please <a href="%s">click here</a> to verify your email address.<br/>
		`, newUser.Username, url),
	}
	sendEmailErr := utils.SendEmail(&user, &emailData)
	if sendEmailErr != nil {
		return errors.New("failed to send email")
	}
	return nil
}

func (i impl) SignUpVerifyEmail(id, code string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	user, getUserErr := i.repo.GetUser(context.Background(), bson.M{"_id": objId})
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
