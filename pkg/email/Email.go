package email

// Email stores the relevant data to send emails.
type Email struct {
	hostname string
	port     string
	address  string
	password string
}

// NewEmail is a helper functon to create a new email.
func (email *Email) NewEmail(host string, port string, addr string, pass string) Email {
	var e Email
	e.hostname = host
	e.port = port
	e.address = addr
	e.password = pass
	return e
}
