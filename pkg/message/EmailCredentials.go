package message

import (
	"github.com/FetchWeb/Fetch/pkg/core"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// EmailCredentials stores the relevant credenial data to send emails.
type EmailCredentials struct {
	core.DBObject
	Address  string `json:"address"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func NewCredentials(addr string, host string, name string, port string, pass string) *EmailCredentials {
	var ec *EmailCredentials
	ec.Address = addr
	ec.Hostname = host
	ec.Name = name
	ec.Port = port
	ec.Password = pass
	return ec
}

func LoadFromConfig(dir string) (*EmailCredentials, error) {
	err := config.Load(file.NewSource(file.WithPath(dir)))
	if err != nil {
		return nil, err
	}

	var ec *EmailCredentials
	err = config.Scan(&ec)
	if err != nil {
		return nil, err
	}

	return ec, nil
}
