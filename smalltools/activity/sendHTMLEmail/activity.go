/*
 * Copyright © 2017. TIBCO Software Inc. [JGrotex]
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
/*
inspired by https://github.com/jpoehls/gophermail
*/
package sendHTMLEmail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"strings"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivServer   = "Server"
	ivPort     = "Port"
	ivSender   = "Sender"
	ivPass     = "Pass"
	ivTo       = "To"
	ivSubject  = "Subject"
	ivHTML     = "HTML"
	ovfeedback = "feedback"

	maxLength = 76
	crlf      = "\r\n"
)

type Message struct {
	From     mail.Address
	ReplyTo  mail.Address
	To       []mail.Address
	Subject  string
	HTMLBody string
	Headers  mail.Header
}

type splittingWriter struct {
	b       *bytes.Buffer
	w       io.Writer
	flushed bool
}

var (
	activityLog = logger.GetLogger("tools-activity-emailvalidation")
	delimiter   = []byte("\r\n")

	ErrMissingRecipient   = errors.New("No recipient specified. Please specify To recipient.")
	ErrMissingFromAddress = errors.New("No from address specified.")
)

type EmailValidationActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &EmailValidationActivity{metadata: metadata}
}

func (a *EmailValidationActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *EmailValidationActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing sendHTMLEmail activity")
	//Read Inputs
	if context.GetInput(ivServer) == nil {
		return false, activity.NewError("Server string is not configured", "sendHTMLEmail-4001", nil)
	}
	server := context.GetInput(ivServer).(string)
	if context.GetInput(ivPort) == nil {
		return false, activity.NewError("Server Port string is not configured", "sendHTMLEmail-4002", nil)
	}
	port := context.GetInput(ivPort).(string)
	if context.GetInput(ivSender) == nil {
		return false, activity.NewError("Email Sender string is not configured", "sendHTMLEmail-4003", nil)
	}
	sender := context.GetInput(ivSender).(string)
	if context.GetInput(ivPass) == nil {
		return false, activity.NewError("Password string is not configured", "sendHTMLEmail-4004", nil)
	}
	pass := context.GetInput(ivPass).(string)
	if context.GetInput(ivTo) == nil {
		return false, activity.NewError("TO Addr string is not configured", "sendHTMLEmail-4005", nil)
	}
	to := context.GetInput(ivTo).(string)
	if context.GetInput(ivSubject) == nil {
		return false, activity.NewError("Subject string is not configured", "sendHTMLEmail-4006", nil)
	}
	subject := context.GetInput(ivSubject).(string)
	if context.GetInput(ivHTML) == nil {
		return false, activity.NewError("Body HTML string is not configured", "sendHTMLEmail-4007", nil)
	}
	htmlbody := context.GetInput(ivHTML).(string)

	//sendHTMLEMail - Start
	activityLog.Info("Executing sendHTMLEmail with: " + server + ":" + port)

	tlsconfig := tls.Config{
		InsecureSkipVerify: true,
		ServerName:         server,
	}
	auth := smtp.PlainAuth(
		"",
		sender,
		pass,
		server,
	)

	msg := &Message{}
	msg.SetFrom(sender)
	msg.SetReplyTo(sender)
	msg.AddTo(to)
	msg.Subject = subject
	msg.HTMLBody = htmlbody
	msg.Headers = mail.Header{}
	msg.Headers["Date"] = []string{time.Now().UTC().Format(time.RFC822)}

	err = SendHTMLMail(sender, server, port, auth, msg, tlsconfig)
	if err != nil {
		return false, activity.NewError("SendHTMLMail Error during Execution", "sendHTMLEmail-5001", err)
	}

	context.SetOutput(ovfeedback, "done")
	//createHTML - End

	return true, nil
}

//SendHTMLMail send TLS HTML Email - Main Function
func SendHTMLMail(sender string, server string, port string, a smtp.Auth, msg *Message, cfg tls.Config) error {
	addr := server + ":" + port
	msgBytes, err := msg.Bytes()
	var to []string
	for _, address := range msg.To {
		to = append(to, address.Address)
	}
	activityLog.Info("... sendHTMLEmail - dial with: " + addr)
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	activityLog.Info("... sendHTMLEmail - startTLS with")
	if ok, _ := c.Extension("STARTTLS"); ok {
		if err = c.StartTLS(&cfg); err != nil {
			return err
		}
	}
	activityLog.Info("... sendHTMLEmail - AUTH")
	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				return err
			}
		}
	}
	activityLog.Info("... sendHTMLEmail - Mail with " + sender)
	if err = c.Mail(sender); err != nil {
		return err
	}
	activityLog.Info("... sendHTMLEmail - Rcpt " + addr)
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	activityLog.Info("... sendHTMLEmail - Data")
	w, err := c.Data()
	if err != nil {
		return err
	}

	activityLog.Info("... sendHTMLEmail - Write")
	_, err = w.Write(msgBytes)
	if err != nil {
		return err
	}

	activityLog.Info("... sendHTMLEmail - Close")
	err = w.Close()
	if err != nil {
		return err
	}
	activityLog.Info("... sendHTMLEmail - Quit")
	return c.Quit()
}

//SendHTMLMail send TLS HTML Email - Main Function - END
// ***

func getAddressListString(addresses []mail.Address) string {
	var addressStrings []string

	for _, address := range addresses {
		addressStrings = append(addressStrings, address.String())
	}
	return strings.Join(addressStrings, ","+crlf+" ")
}

func appendMailAddresses(dest *[]mail.Address, addresses ...string) error {
	var parsedAddresses []mail.Address
	var err error

	for _, address := range addresses {
		parsed, err := mail.ParseAddress(address)
		if err != nil {
			return err
		}
		parsedAddresses = append(parsedAddresses, *parsed)
	}

	*dest = append(*dest, parsedAddresses...)
	return err
}

func setMailAddress(dest *mail.Address, address string) error {
	parsed, err := mail.ParseAddress(address)
	if err != nil {
		return err
	}
	*dest = *parsed
	return nil
}

//*** getters and setters

func (m *Message) SetFrom(address string) error {
	return setMailAddress(&m.From, address)
}
func (m *Message) SetReplyTo(address string) error {
	return setMailAddress(&m.ReplyTo, address)
}
func (m *Message) AddTo(addresses ...string) error {
	return appendMailAddresses(&m.To, addresses...)
}
func (m *Message) Bytes() ([]byte, error) {
	var buffer = &bytes.Buffer{}
	header := textproto.MIMEHeader{}
	return m.bytes(buffer, header)
}

//*** create Headers

func (m *Message) bytes(buffer *bytes.Buffer, header textproto.MIMEHeader) ([]byte, error) {
	var err error

	toAddrs := getAddressListString(m.To)
	var hasTo = toAddrs != ""

	if !hasTo {
		return nil, ErrMissingRecipient
	}
	if hasTo {
		header.Add("To", toAddrs)
	}

	var emptyAddress mail.Address
	if m.From == emptyAddress {
		return nil, ErrMissingFromAddress
	}
	header.Add("From", m.From.String())
	if m.ReplyTo != emptyAddress {
		header.Add("Reply-To", m.ReplyTo.String())
	}
	if m.Subject != "" {
		header.Add("Subject", m.Subject)
	}
	for k, v := range m.Headers {
		header[k] = v
	}

	multipartw := multipart.NewWriter(buffer)
	header.Add("MIME-Version", "1.0")
	header.Add("Content-Type", fmt.Sprintf("multipart/mixed;%s boundary=%s", crlf, multipartw.Boundary()))

	err = writeHeader(buffer, header)
	if err != nil {
		return nil, err
	}

	_, err = fmt.Fprintf(buffer, "--%s%s", multipartw.Boundary(), crlf)
	if err != nil {
		return nil, err
	}

	if m.HTMLBody != "" {

		altw := multipart.NewWriter(buffer)
		header = textproto.MIMEHeader{}
		header.Add("Content-Type", fmt.Sprintf("multipart/alternative;%s boundary=%s", crlf, altw.Boundary()))
		err := writeHeader(buffer, header)
		if err != nil {
			return nil, err
		}

		if m.HTMLBody != "" {
			header = textproto.MIMEHeader{}
			header.Add("Content-Type", "text/html; charset=utf-8")
			header.Add("Content-Transfer-Encoding", "base64")

			partw, err := altw.CreatePart(header)
			if err != nil {
				return nil, err
			}

			htmlBodyBytes := []byte(m.HTMLBody)
			encoder := NewBase64MimeEncoder(partw)
			_, err = encoder.Write(htmlBodyBytes)
			if err != nil {
				return nil, err
			}
			err = encoder.Close()
			if err != nil {
				return nil, err
			}
		}
		altw.Close()
	}

	multipartw.Close()
	return buffer.Bytes(), nil
}

//*** Write Header helper

func writeHeader(w io.Writer, header textproto.MIMEHeader) error {
	for k, vs := range header {
		_, err := fmt.Fprintf(w, "%s: ", k)
		if err != nil {
			return err
		}

		for i, v := range vs {
			v = textproto.TrimString(v)

			_, err := fmt.Fprintf(w, "%s", v)
			if err != nil {
				return err
			}

			if i < len(vs)-1 {
				return errors.New("Multiple header values are not supported.")
			}
		}

		_, err = fmt.Fprint(w, crlf)
		if err != nil {
			return err
		}
	}

	_, err := fmt.Fprint(w, crlf)
	if err != nil {
		return err
	}

	return nil
}

//*** Minetype helpers

func (t *splittingWriter) Write(p []byte) (n int, err error) {
	n, err = t.b.Write(p)
	for t.b.Len() >= maxLength {
		if t.flushed {
			_, err = t.w.Write(delimiter)
			if err != nil {
				return
			}
		}
		var n2 int64
		n2, err = io.CopyN(t.w, t.b, maxLength)
		if n2 > 0 {
			t.flushed = true
		}
	}
	return
}

func (t *splittingWriter) Close() (err error) {
	if t.b.Len() > 0 {
		if t.flushed {
			_, err = t.w.Write(delimiter)
			if err != nil {
				return
			}
		}
		var n int64
		n, err = io.Copy(t.w, t.b)
		if n > 0 {
			t.flushed = true
		}
	}
	return
}

type base64MimeEncoder struct {
	enc io.WriteCloser
	w   io.WriteCloser
}

func (t *base64MimeEncoder) Write(p []byte) (n int, err error) {
	n, err = t.enc.Write(p)
	return
}

func (t *base64MimeEncoder) Close() (err error) {
	err = t.enc.Close()
	if err != nil {
		return err
	}
	err = t.w.Close()
	return
}

func NewBase64MimeEncoder(w io.Writer) io.WriteCloser {
	splitter := &splittingWriter{
		w: w,
		b: &bytes.Buffer{},
	}
	t := &base64MimeEncoder{w: splitter}
	t.enc = base64.NewEncoder(base64.StdEncoding, splitter)
	return t
}
