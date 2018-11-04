package message

import (
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// EmailCredentials stores the relevant credenial data to send emails.
type EmailCredentials struct {
	Address  string `json:"address"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

// LoadFromConfig loads the email credentials from a JSON config file.
func (ec *EmailCredentials) LoadFromConfig(dir string) error {
	err := config.Load(file.NewSource(file.WithPath(dir)))
	if err != nil {
		return err
	}

	err = config.Scan(&ec)
	if err != nil {
		return err
	}

	return nil
}
