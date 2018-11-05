package main

import (
	"io/ioutil"
	"log"
	"net/mail"

	"github.com/FetchWeb/Fetch/pkg/message"
)

func main() {
	var creds message.EmailCredentials
	if err := creds.LoadFromConfig("../../configs/TestEmailConfig.json"); err != nil {
		log.Fatal(err)
	}

	buff, err := ioutil.ReadFile("../../test/message/TestEmailTemplate.html")
	if err != nil {
		log.Fatal(err)
	}

	data := message.NewHTMLMessage("Queue Test Email: ", string(buff))
	data.From = mail.Address{Name: creds.Name, Address: creds.Address}
	data.To = []string{"taliesinwrmillhouse@gmail.com"}

	email := &message.Email{
		Credentials: &creds,
		Data:        data,
	}

	var s message.Service
	if err := s.Startup(); err != nil {
		panic(err)
	}

	if err := s.EnqueueEmail(email); err != nil {
		panic(err)
	}

	if err := s.DequeueEmail(); err != nil {
		panic(err)
	}
}
