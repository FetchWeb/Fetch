package email

import (
	"net/smtp"
)

// MailService handles emails and other simlar functionality within the package.
type MailService struct {
}

// SendEmail sends an email with the message to the recipients from the sender.
func (mailService *MailService) SendEmail(sender *Email, recipients []string, message string) error {
	return smtp.SendMail(sender.Hostname+":"+sender.Port, smtp.PlainAuth("", sender.Address, sender.Password, sender.Hostname), sender.Address, recipients, []byte(message))
}
