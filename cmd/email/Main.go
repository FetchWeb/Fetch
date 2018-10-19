package main

import (
	"fmt"
	"github.com/FetchWeb/Fetch/pkg/email"
)

func main() {
	e, err := email.LoadFromConfig("../../configs/TestEmailConfig.json")
	if err != nil {
		panic(err)
	}
	fmt.Print(e)

	var mailService email.MailService
	err = mailService.SendEmail(e, []string{"taliesinwrmillhouse@gmail.com"}, "Hello world!")
	if err != nil {
		panic(err)
	}
}
