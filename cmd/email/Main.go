package main

import (
	"io/ioutil"
	"log"
	"net/mail"

	"github.com/FetchWeb/Fetch/pkg/email"
)

func main() {
	// Load in config.
	creds, err := email.LoadFromConfig("../../configs/TestEmailConfig.json")
	if err != nil {
		log.Fatal(err)
	}

	// Load in template.
	buff, err := ioutil.ReadFile("../../test/email/TestEmailTemplate.html")
	if err != nil {
		log.Fatal(err)
	}

	// Create & initialize message.
	m := email.NewHTMLMessage("This is a subject", string(buff))
	m.From = mail.Address{Name: creds.Name, Address: creds.Address}
	m.To = []string{"taliesinwrmillhouse@gmail.com"}

	// Add attachment.
	if err := m.AddAttachment("../../test/email/attachment.jpg", false); err != nil {
		log.Fatal(err)
	}

	// Add header.
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// Send.
	var emailService email.Service
	err = emailService.SendEmail(creds, m)
	if err != nil {
		log.Fatal(err)
	}
}
