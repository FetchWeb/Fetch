package message

import (
	"github.com/FetchWeb/Fetch/pkg/message"
	"net/smtp"
	"sync"

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
	b, err := e.MarshalBinary()
	if err != nil {
		return err
	}

	key := core.UniqueID()
	s.clientKeys.Push(key)

	err = s.client.Set(key, b, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SendEmails() error {
	while (s.clientKeys.CanPop()) {
		
	}
	if !s.clientKeys.CanPop() {
		return nil;
	}

	val, err := s.client.Get(s.clientKeys.Pop().(string)).Result()
	if err != nil {
		return err
	}

	var e message.Email
}

// SendEmail sends an email with the message to the recipients from the sender.
func (s *Service) SendEmail(ec *EmailCredentials, ed *EmailData) error {
	return smtp.SendMail(ec.Hostname+":"+ec.Port, smtp.PlainAuth("", ec.Address, ec.Password, ec.Hostname), ec.Address, ed.GetRecipients(), ed.Data())
}

// SendEmailAsync sends an email using a goroutine with the message to the recipients from the sender.
func (s *Service) SendEmailAsync(ec *EmailCredentials, ed *EmailData, wg *sync.WaitGroup) error {
	defer wg.Done()
	return smtp.SendMail(ec.Hostname+":"+ec.Port, smtp.PlainAuth("", ec.Address, ec.Password, ec.Hostname), ec.Address, ed.GetRecipients(), ed.Data())
}
