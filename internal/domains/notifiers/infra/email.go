package infra

import (
	"fmt"
	"log"
	"pantori/internal/domains/notifiers/core"

	"net/smtp"
)

type email struct{}

func NewEmailProvider() *email {
	return &email{}
}

func (e *email) SendEmail(user core.User, expireToday, expireSoon, expired []core.Good) error {
	if len(expireToday) == 0 {
		expireToday = append(expireToday, core.Good{Name: "empty"})
	}
	if len(expireSoon) == 0 {
		expireSoon = append(expireSoon, core.Good{Name: "empty"})
	}
	if len(expired) == 0 {
		expired = append(expired, core.Good{Name: "empty"})
	}
	fmt.Printf("%s -- today: %s --- soon: %s --- expired: %s\n", user.Name, expireToday[0].Name, expireSoon[0].Name, expired[0].Name)
	return nil
}

func foo() {

	// Choose auth method and set it up

	auth := smtp.PlainAuth("", "john.doe@gmail.com", "extremely_secret_pass", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	to := []string{"kate.doe@example.com"}

	msg := []byte("To: kate.doe@example.com\r\n" +

		"Subject: Why aren't you using Mailtrap yet?\r\n" +

		"\r\n" +

		"Here's the space for our great sales pitch\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "john.doe@gmail.com", to, msg)

	if err != nil {

		log.Fatal(err)

	}

}
