package libs

import (
	"net/smtp"
	"fmt"
)

func Mailer(to string, subject string, message string){
	user := "b359164a-fbb9-416b-829d-0802dc67d3cc"
	pass := "b359164a-fbb9-416b-829d-0802dc67d3cc"

	msg := "From: " + "noreply@bitfuture.exchange" + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		message

	err := smtp.SendMail("smtp.postmarkapp.com:587",
		smtp.PlainAuth("", user, pass, "smtp.postmarkapp.com"),
		"noreply@bitfuture.exchange", []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("smtp error: %s", err)
		return
	}
	
	fmt.Println("sent, visit http://foobarbazz.mailinator.com")
}