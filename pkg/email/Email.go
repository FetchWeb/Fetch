package email

import (
	"io/ioutil"
	"path/filepath"

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

type Attachment struct {
	Name string
	Data []byte
}

type Header struct {
	Key   string
	Value string
}

type Message struct {
	To              []string
	CC              []string
	BCC             []string
	ReplyTo         string
	Subject         string
	Body            string
	BodyContentType string
	Headers         []Header
	Attachments     map[string]*Attachment
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

func NewAttachment(name string, data []byte) *Attachment {
	var a *Attachment
	a.Name = name
	a.Data = data
	return a
}

func (m *Message) AddAttachment(dir string) error {
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return err
	}

	_, name := filepath.Split(dir)
	m.Attachments[name] = NewAttachment(name, data)
	
	return nil
}
