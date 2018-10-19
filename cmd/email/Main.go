package main

import (
	"io/ioutil"

	"github.com/FetchWeb/Fetch/pkg/email"
)

func main() {
	e, err := email.LoadFromConfig("../../configs/TestEmailConfig.json")
	if err != nil {
		panic(err)
	}

	messageBuf, err := ioutil.ReadFile("../../test/email/TestEmailTemplate.html")
	if err != nil {
		panic(err)
	}

	var mailService email.MailService
	err = mailService.SendEmail(e, []string{"taliesinwrmillhouse@gmail.com"}, string(messageBuf))
	if err != nil {
		panic(err)
	}
}
