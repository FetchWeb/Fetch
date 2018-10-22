package email

import (
	"github.com/FetchWeb/Fetch/pkg/core"
	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// Credentials stores the relevant credenial data to send emails.
type Credentials struct {
	core.DBObject
	Address  string `json:"address"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func NewCredentials(addr string, host string, name string, port string, pass string) *Credentials {
	var c *Credentials
	c.Address = addr
	c.Hostname = host
	c.Name = name
	c.Port = port
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
