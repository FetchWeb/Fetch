package email

import (
	"net/smtp"
)

// Service handles emails and other simlar functionality within the package.
type Service struct {
}

// SendEmail sends an email with the message to the recipients from the sender.
func (service *Service) SendEmail(sender *Email, recipients []string, message string) error {
	return smtp.SendMail(sender.Hostname+":"+sender.Port, smtp.PlainAuth("", sender.Address, sender.Password, sender.Hostname), sender.Address, recipients, []byte(message))
}
