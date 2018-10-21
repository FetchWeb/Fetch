package main

import (
	"io/ioutil"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/FetchWeb/Fetch/pkg/email"
)

func main() {
	// e, err := email.LoadFromConfig("../../configs/TestEmailConfig.json")
	// if err != nil {
	// 	panic(err)
	// }

	// compose the message
	buff, err := ioutil.ReadFile("../../test/email/TestEmailTemplate.html")
	if err != nil {
		log.Fatal(err)
	}
	m := email.NewHTMLMessage("This is a subject", string(buff))
	m.From = mail.Address{Name: "From", Address: "from@example.com"}
	m.To = []string{"taliesinwrmillhouse@gmail.com"}

	// add attachments
	if err := m.AddAttachment("../../test/email/attachment.jpg", false); err != nil {
		log.Fatal(err)
	}

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", "fetchweb.test@gmail.com", "iumhivtjafwsxhhe", "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		log.Fatal(err)
	}

	var creds email.Credentials
	var emailService email.Service
	emailService.SendEmail()
}
