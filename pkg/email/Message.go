package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"net/mail"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

// Message represents a smtp message.
type Message struct {
	From            mail.Address
	To              []string
	Cc              []string
	Bcc             []string
	ReplyTo         string
	Subject         string
	Body            string
	BodyContentType string
	Headers         []Header
	Attachments     map[string]*Attachment
}

// NewPlainTextMessage returns a new Message that can compose an email with attachments
func NewPlainTextMessage(subject string, body string) *Message {
	m := &Message{Subject: subject, Body: body, BodyContentType: "text/plain"}
	m.Attachments = make(map[string]*Attachment)
	return m
}

// NewHTMLMessage returns a new Message that can compose an HTML email with attachments
func NewHTMLMessage(subject string, body string) *Message {
	m := &Message{Subject: subject, Body: body, BodyContentType: "text/html"}
	m.Attachments = make(map[string]*Attachment)
	return m
}

// AddAttachment adds a new attachment to the message.
func (m *Message) AddAttachment(file string, inline bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(file)

	m.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

// AttachBuffer attaches a binary attachment.
func (m *Message) AttachBuffer(filename string, buf []byte, inline bool) error {
	m.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     buf,
		Inline:   inline,
	}
	return nil
}

// AddHeader ads a Header to message
func (m *Message) AddHeader(key string, value string) Header {
	newHeader := Header{Key: key, Value: value}
	m.Headers = append(m.Headers, newHeader)
	return newHeader
}

// GetRecipients returns all the recipients of the message.
func (m *Message) GetRecipients() []string {
	recipients := m.To

	for _, cc := range m.Cc {
		recipients = append(recipients, cc)
	}

	for _, bcc := range m.Bcc {
		recipients = append(recipients, bcc)
	}

	return recipients
}

// Data returns all the message data as a byte array.
func (m *Message) Data() []byte {
	buf := bytes.NewBuffer(nil)

	// Write from and recipients.
	buf.WriteString("From: " + m.From.String() + "\n")
	t := time.Now()
	buf.WriteString("Date: " + t.Format(time.RFC1123Z) + "\n")
	buf.WriteString("To: " + strings.Join(m.To, ",") + "\n")
	if len(m.Cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(m.Cc, ",") + "\n")
	}

	// Write encoding.
	var coder = base64.StdEncoding
	var subject = "=?UTF-8?B?" + coder.EncodeToString([]byte(m.Subject)) + "?="
	buf.WriteString("Subject: " + subject + "\n")

	if len(m.ReplyTo) > 0 {
		buf.WriteString("Reply-To: " + m.ReplyTo + "\n")
	}

	buf.WriteString("MIME-Version: 1.0\n")

	// Write headers.
	if len(m.Headers) > 0 {
		for _, header := range m.Headers {
			buf.WriteString(fmt.Sprintf("%s: %s\n", header.Key, header.Value))
		}
	}

	// Write attachments.
	boundary := "f46d043c813270fc6b04c2d223da"
	if len(m.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\n")
		buf.WriteString("\n--" + boundary + "\n")
	}

	// Write content type.
	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n\n", m.BodyContentType))
	buf.WriteString(m.Body)
	buf.WriteString("\n")

	if len(m.Attachments) > 0 {
		for _, attachment := range m.Attachments {
			buf.WriteString("\n\n--" + boundary + "\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: message/rfc822\n")
				buf.WriteString("Content-Disposition: inline; filename=\"" + attachment.Filename + "\"\n\n")

				buf.Write(attachment.Data)
			} else {
				ext := filepath.Ext(attachment.Filename)
				mimetype := mime.TypeByExtension(ext)
				if mimetype != "" {
					mime := fmt.Sprintf("Content-Type: %s\n", mimetype)
					buf.WriteString(mime)
				} else {
					buf.WriteString("Content-Type: application/octet-stream\n")
				}
				buf.WriteString("Content-Transfer-Encoding: base64\n")

				buf.WriteString("Content-Disposition: attachment; filename=\"=?UTF-8?B?")
				buf.WriteString(coder.EncodeToString([]byte(attachment.Filename)))
				buf.WriteString("?=\"\n\n")

				b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
				base64.StdEncoding.Encode(b, attachment.Data)

				for i, l := 0, len(b); i < l; i++ {
					buf.WriteByte(b[i])
					if (i+1)%76 == 0 {
						buf.WriteString("\n")
					}
				}
			}

			buf.WriteString("\n--" + boundary)
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

// Send sends the message.
func Send(addr string, auth smtp.Auth, m *Message) error {
	return smtp.SendMail(addr, auth, m.From.Address, m.GetRecipients(), m.Data())
}
