package model

type User struct {
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}
