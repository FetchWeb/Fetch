package message

import (
	"net/smtp"
	"strings"

	"github.com/FetchWeb/Fetch/pkg/core"
	"github.com/go-redis/redis"
)

// Service handles emails and other simlar functionality within the package.
type Service struct {
	client     *redis.Client
	clientKeys core.Queue
}

func (s *Service) Startup() error {
	s.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := s.client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) EnqueueEmail(e *Email) error {
	data, err := e.MarshalBinary()
	if err != nil {
		return err
	}

	key := core.UniqueID()
	s.clientKeys.Push(key)

	err = s.client.Set(key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DequeueEmail() error {
	if !s.clientKeys.CanPop() {
		return nil
	}

	key := s.clientKeys.Pop().(string)

	data, err := s.client.Get(key).Result()
	if err != nil {
		return err
	}

	var email Email
	err = email.UnmarshalBinary([]byte(data))
	if err != nil {
		return err
	}

	addr := strings.Join([]string{email.Credentials.Hostname, ":", email.Credentials.Port}, "")
	auth := smtp.PlainAuth("", email.Credentials.Address, email.Credentials.Password, email.Credentials.Hostname)

	err = smtp.SendMail(addr, auth, email.Credentials.Address, email.Data.GetRecipients(), email.Data.Data())
	if err != nil {
		return err
	}

	_, err = s.client.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}

// // SendEmail sends an email with the message to the recipients from the sender.
// func (s *Service) SendEmail(ec *EmailCredentials, ed *EmailData) error {
// 	return smtp.SendMail(ec.Hostname+":"+ec.Port, smtp.PlainAuth("", ec.Address, ec.Password, ec.Hostname), ec.Address, ed.GetRecipients(), ed.Data())
// }

// // SendEmailAsync sends an email using a goroutine with the message to the recipients from the sender.
// func (s *Service) SendEmailAsync(ec *EmailCredentials, ed *EmailData, wg *sync.WaitGroup) error {
// 	defer wg.Done()
// 	return smtp.SendMail(ec.Hostname+":"+ec.Port, smtp.PlainAuth("", ec.Address, ec.Password, ec.Hostname), ec.Address, ed.GetRecipients(), ed.Data())
// }
