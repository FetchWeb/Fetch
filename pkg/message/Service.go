package message

import (
	"net/smtp"
	"sync"
)

// Service handles emails and other simlar functionality within the package.
type Service struct {
}

// SendEmail sends an email with the message to the recipients from the sender.
func (service *Service) SendEmail(ec *EmailCredentials, ed *EmailData) error {
	return smtp.SendMail(ec.Hostname+":"+ec.Port, smtp.PlainAuth("", ec.Address, ec.Password, ec.Hostname), ec.Address, ed.GetRecipients(), ed.Data())
}

// SendEmailAsync sends an email using a goroutine with the message to the recipients from the sender.
func (service *Service) SendEmailAsync(ec *EmailCredentials, ed *EmailData, wg *sync.WaitGroup) error {
	defer wg.Done()
	return smtp.SendMail(ec.Hostname+":"+ec.Port, smtp.PlainAuth("", ec.Address, ec.Password, ec.Hostname), ec.Address, ed.GetRecipients(), ed.Data())
}
