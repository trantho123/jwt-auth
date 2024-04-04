package utils

import (
	"Jwtwithecdsa/api/internal/model"
	"errors"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"gopkg.in/gomail.v2"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

type EmailData struct {
	URL     string
	Subject string
	Content string
}

func SendEmail(user *model.User, emailData *EmailData) error {

	email := gomail.NewMessage()
	email.SetHeader("From", os.Getenv("EMAIL_FROM"))
	email.SetHeader("To", user.Email)
	email.SetHeader("Subject", emailData.Subject)
	email.SetBody("text/html", emailData.Content)
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))
	if err := d.DialAndSend(email); err != nil {
		return err
	}

	return nil
}

func IsValidEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email address")
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
