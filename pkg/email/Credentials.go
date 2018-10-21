package email

import (
	"github.com/FetchWeb/Fetch/pkg/core"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// Credentials stores the relevant credenial data to send emails.
type Credentials struct {
	core.DBObject
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func NewCredentials(host string, port string, addr string, pass string) *Credentials {
	var c *Credentials
	c.Hostname = host
	c.Port = port
	c.Address = addr
	c.Password = pass
	return c
}

func LoadFromConfig(dir string) (*Credentials, error) {
	err := config.Load(file.NewSource(file.WithPath(dir)))
	if err != nil {
		return nil, err
	}

	var c *Credentials
	err = config.Scan(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
