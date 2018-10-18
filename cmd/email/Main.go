package main

import "github.com/FetchWeb/Fetch/pkg/email"

func main() {
	var mailService email.MailService
	mailService.SendEmail()
}
