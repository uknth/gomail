// Package gomail is wrapper around smtp.SendMaile
package gomail

import (
	"errors"
	"net/smtp"
)

var errSendingMail = errors.New("Error sending email")

// EmailSender defines interface to be used for mailinng
// It gives an abstraction for test cases to mock the data
type Mailer interface {
	Send(subject, content string) error
}

// Config config object for email
type Config struct {
	UserName   string
	Password   string
	ServerHost string
	ServerPort string
	SenderAddr string
	Prefix     string
	Receivers  []string // Default Receivers
}

// Sender implements EmailSender interface and is used to send email
type DefaultMailer struct {
	cf   *Config
	send func(string, smtp.Auth, string, []string, []byte) error
}

// Send is called to send e-mail
func (m *DefaultMailer) Send(subject, content string) error {
	plainAuth := smtp.PlainAuth("",
		m.cf.UserName,
		m.cf.Password,
		m.cf.ServerHost,
	)

	to := m.cf.Receivers

	subject = m.cf.Prefix + " :: " + subject
	message := "Subject: " + subject
	message += "\r\n" + content

	addr := m.cf.ServerHost + ":" + m.cf.ServerPort

	if err := m.send(
		addr,
		plainAuth,
		m.cf.SenderAddr,
		to,
		[]byte(message),
	); err != nil {
		return err
	}

	return nil
}

// New initializes the mailer
func New(config *Config) Mailer {
	return &DefaultMailer{
		cf: config,
		send: func(
			addr string,
			plainAuth smtp.Auth,
			senderAddr string,
			addresses []string,
			message []byte,
		) error {
			err := smtp.SendMail(
				addr,
				plainAuth,
				senderAddr,
				addresses,
				message,
			)
			if err != nil {
				return errors.New("Unable to send mail: " + err.Error())
			}
			return nil
		},
	}
}
