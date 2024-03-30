package utils

import "gopkg.in/gomail.v2"

func SendEmail(to, subject, body string) error {

	email := gomail.NewMessage()
	email.SetHeader("From", "trandinhttho@gmail.com")
	email.SetHeader("To", to)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, "trandinhttho@gmail.com", "azbw hxfv qtvg fzlq")
	if err := d.DialAndSend(email); err != nil {
		return err
	}

	return nil
}
