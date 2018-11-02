package message

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

// EmailData represents a smtp message.
type EmailData struct {
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

// NewPlainTextMessage returns a new EmailData that can compose an email with attachments
func NewPlainTextMessage(subject string, body string) *EmailData {
	ed := &EmailData{Subject: subject, Body: body, BodyContentType: "text/plain"}
	ed.Attachments = make(map[string]*Attachment)
	return ed
}

// NewHTMLMessage returns a new EmailData that can compose an HTML email with attachments
func NewHTMLMessage(subject string, body string) *EmailData {
	ed := &EmailData{Subject: subject, Body: body, BodyContentType: "text/html"}
	ed.Attachments = make(map[string]*Attachment)
	return ed
}

// AddAttachment adds a new attachment to the message.
func (ed *EmailData) AddAttachment(file string, inline bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(file)

	ed.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

// AttachBuffer attaches a binary attachment.
func (ed *EmailData) AttachBuffer(filename string, buf []byte, inline bool) error {
	ed.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     buf,
		Inline:   inline,
	}
	return nil
}

// AddHeader ads a Header to message
func (ed *EmailData) AddHeader(key string, value string) Header {
	newHeader := Header{Key: key, Value: value}
	ed.Headers = append(ed.Headers, newHeader)
	return newHeader
}

// GetRecipients returns all the recipients of the message.
func (ed *EmailData) GetRecipients() []string {
	recipients := ed.To

	for _, cc := range ed.Cc {
		recipients = append(recipients, cc)
	}

	for _, bcc := range ed.Bcc {
		recipients = append(recipients, bcc)
	}

	return recipients
}

// Data returns all the message data as a byte array.
func (ed *EmailData) Data() []byte {
	buf := bytes.NewBuffer(nil)

	// Write from and recipients.
	buf.WriteString("From: " + ed.From.String() + "\n")
	t := time.Now()
	buf.WriteString("Date: " + t.Format(time.RFC1123Z) + "\n")
	buf.WriteString("To: " + strings.Join(ed.To, ",") + "\n")
	if len(ed.Cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(ed.Cc, ",") + "\n")
	}

	// Write encoding.
	var coder = base64.StdEncoding
	var subject = "=?UTF-8?B?" + coder.EncodeToString([]byte(ed.Subject)) + "?="
	buf.WriteString("Subject: " + subject + "\n")

	if len(ed.ReplyTo) > 0 {
		buf.WriteString("Reply-To: " + ed.ReplyTo + "\n")
	}

	buf.WriteString("MIME-Version: 1.0\n")

	// Write headers.
	if len(ed.Headers) > 0 {
		for _, header := range ed.Headers {
			buf.WriteString(fmt.Sprintf("%s: %s\n", header.Key, header.Value))
		}
	}

	// Write boundary.
	boundary := "f46d043c813270fc6b04c2d223da"
	if len(ed.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\n")
		buf.WriteString("\n--" + boundary + "\n")
	}

	// Write content type.
	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n\n", ed.BodyContentType))
	buf.WriteString(ed.Body)
	buf.WriteString("\n")

	// Write attachments.
	if len(ed.Attachments) > 0 {
		for _, attachment := range ed.Attachments {
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
func Send(addr string, auth smtp.Auth, ed *EmailData) error {
	return smtp.SendMail(addr, auth, ed.From.Address, ed.GetRecipients(), ed.Data())
}
