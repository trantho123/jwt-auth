package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Email            string             `bson:"email"`
	Username         string             `bson:"username"`
	Password         string             `bson:"password"`
	VerificationCode string
	Verified         bool      `bson:"verified"`
	CreatedAt        time.Time `bson:"createdat"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type VerifyOTP struct {
	Email   string `json:"email"`
	OTPCode string `json:"otpcode"`
}

type LoginInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Email     string             `json:"email"`
	Username  string             `json:"username"`
	CreatedAt time.Time          `json:"createdAt"`
}
