package email

import (
	"net/smtp"
)

// MailService handles emails and other simlar functionality within the package.
type MailService struct {
	emails []Email
}

// SendEmail sends an email with the message to the recipients from the sender.
func (mailService *MailService) SendEmail(sender *Email, recipients []string, message string) error {
	return smtp.SendMail(sender.hostname+":"+sender.port, smtp.PlainAuth("", sender.address, sender.password, sender.hostname), sender.address, recipients, []byte(message))
}
