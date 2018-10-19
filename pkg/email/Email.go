package email

import (
	"github.com/FetchWeb/Fetch/pkg/core"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// Email stores the relevant data to send emails.
type Email struct {
	core.DBObject
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

// NewEmail is a helper functon to create a new email.
func NewEmail(host string, port string, addr string, pass string) *Email {
	var e *Email
	e.Hostname = host
	e.Port = port
	e.Address = addr
	e.Password = pass
	return e
}

// LoadFromConfig is a helper function to load the email from a config file.
func LoadFromConfig(configDir string) (*Email, error) {
	err := config.Load(file.NewSource(file.WithPath(configDir)))
	if err != nil {
		return nil, err
	}

	var e *Email
	err = config.Scan(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
