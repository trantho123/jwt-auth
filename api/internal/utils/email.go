package utils

import (
	"Jwtwithecdsa/api/internal/model"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

type EmailData struct {
	URL      string
	UserName string
	Subject  string
}

func SendEmail(user *model.User, emailData *EmailData) error {

	email := gomail.NewMessage()
	email.SetHeader("From", os.Getenv("EMAIL_FROM"))
	email.SetHeader("To", user.Email)
	email.SetHeader("Subject", emailData.Subject)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, emailData.UserName, emailData.URL)
	email.SetBody("text/html", content)
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
	if err := d.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
