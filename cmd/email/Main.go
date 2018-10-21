package main

import (
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

	// // messageBuf, err := ioutil.ReadFile("../../test/email/TestEmailTemplate.html")
	// // if err != nil {
	// // 	panic(err)
	// // }

	// var service email.Service
	// err = service.SendEmail(e, []string{"taliesinwrmillhouse@gmail.com"}, string("Hello world!"))
	// if err != nil {
	// 	panic(err)
	// }

	// compose the message
	m := email.NewMessage("Hi", "this is the body")
	m.From = mail.Address{Name: "From", Address: "from@example.com"}
	m.To = []string{"taliesinwrmillhouse@gmail.com"}

	// add attachments
	if err := m.Attach("../../test/email/attachment.jpg", true); err != nil {
		log.Fatal(err)
	}

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", "fetchweb.test@gmail.com", "iumhivtjafwsxhhe", "smtp.gmail.com")
	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		log.Fatal(err)
	}
}
