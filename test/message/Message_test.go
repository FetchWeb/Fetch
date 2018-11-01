package message_test

import (
	"io/ioutil"
	"log"
	"net/mail"
	"sync"
	"testing"

	"github.com/FetchWeb/Fetch/pkg/core"
	"github.com/FetchWeb/Fetch/pkg/message"
)

func BenchmarkSendEmailQueueWithGoroutine(b *testing.B) {
	emailCreds, err := message.LoadFromConfig("../../configs/TestEmailConfig.json")
	if err != nil {
		log.Fatal(err)
	}

	buff, err := ioutil.ReadFile("../../test/message/TestEmailTemplate.html")
	if err != nil {
		log.Fatal(err)
	}

	var queue core.Queue
	for i := 0; i < 3; i++ {
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
	for j := 0; j < 3; j++ {
		if queue.CanPop() {
			email := queue.Pop().(*message.Email)
			if err = serivce.SendEmail(email.Credentials, email.Data); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func BenchmarkSendQueueWithoutGoroutine(b *testing.B) {
	emailCreds, err := message.LoadFromConfig("../../configs/TestEmailConfig.json")
	if err != nil {
		log.Fatal(err)
	}

	buff, err := ioutil.ReadFile("../../test/message/TestEmailTemplate.html")
	if err != nil {
		log.Fatal(err)
	}

	var queue core.Queue
	for i := 0; i < 3; i++ {
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
	var wg sync.WaitGroup
	for j := 0; j < 3; j++ {
		if queue.CanPop() {
			email := queue.Pop().(*message.Email)
			wg.Add(1)
			go func() {
				if err = serivce.SendEmailAsync(email.Credentials, email.Data, &wg); err != nil {
					log.Fatal(err)
				}
			}()
		}
	}
	wg.Wait()
}
