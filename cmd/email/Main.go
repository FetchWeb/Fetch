package main

import (
	"io/ioutil"
	"log"
	"net/mail"

	"github.com/FetchWeb/Fetch/pkg/core"
	"github.com/FetchWeb/Fetch/pkg/message"
)

func main() {
	emailCreds, err := message.LoadFromConfig("../../configs/TestEmailConfig.json")
	if err != nil {
		log.Fatal(err)
	}

	buff, err := ioutil.ReadFile("../../test/message/TestEmailTemplate.html")
	if err != nil {
		log.Fatal(err)
	}

	var queue core.Queue
	for i := 0; i < 10; i++ {
		emailData := message.NewHTMLMessage("Queue Test Email: "+string(i), string(buff))
		emailData.From = mail.Address{Name: emailCreds.Name, Address: emailCreds.Address}
		emailData.To = []string{"taliesinwrmillhouse@gmail.com"}

		if err := emailData.AddAttachment("../../test/message/attachment.jpg", false); err != nil {
			log.Fatal(err)
		}

		email := &message.Email{Data: emailData, Credentials: emailCreds}

		queue.Push(email)
	}

	var serivce message.Service
	for i := 0; i > 10; i-- {
		if queue.CanPop() {
			email := queue.Pop().(*message.Email)
			serivce.SendEmail(email.Credentials, email.Data)
		}
	}
}
