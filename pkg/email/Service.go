package email

import (
	"net/smtp"
)

// Service handles emails and other simlar functionality within the package.
type Service struct {
}

// SendEmail sends an email with the message to the recipients from the sender.
func (service *Service) SendEmail(c *Credentials, m *Message) error {
	return smtp.SendMail(c.Hostname+":"+c.Port, smtp.PlainAuth("", c.Address, c.Password, c.Hostname), c.Address, m.GetRecipients(), m.Data())
}
