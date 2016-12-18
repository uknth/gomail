package gomail

import (
	"net/smtp"

	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type emailRecorder struct {
	addr    string
	auth    smtp.Auth
	from    string
	to      []string
	message []byte
}

func mockSend(expectedError error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
	recorder := new(emailRecorder)
	return func(addr string, auth smtp.Auth, from string, to []string, message []byte) error {
		*recorder = emailRecorder{
			addr, auth, from, to, message,
		}
		return expectedError
	}, recorder
}

func mockedMessage(appName, subject, content string) []byte {
	return []byte("Subject: " +
		appName + " :: " + subject + "\r\n" + content)
}

func sender(f func(string, smtp.Auth, string, []string, []byte) error) Mailer {
	return &DefaultMailer{
		cf: &Config{
			UserName:   "user@unbxd.com",
			Password:   "password",
			ServerHost: "mail.unbxd.com",
			ServerPort: "447",
			SenderAddr: "user@unbxd.com",
			Prefix:     "testNode",
			Receivers:  []string{"reciever@unbxd.com"},
		},
		send: f,
	}
}

func TestEmailSend(t *testing.T) {
	Convey("Send Mail", t, func() {
		Convey("Send mocked mail", func() {
			f, rec := mockSend(nil)
			mailSender := sender(f)
			mailSender.Send("hello subject", "hello content")
			expectedMessage := mockedMessage("testNode", "hello subject", "hello content")
			So(string(rec.message), ShouldEqual, string(expectedMessage))
		})
	})
}

func TestFailureCase(t *testing.T) {
	Convey("Send Mail Failure", t, func() {
		Convey("Send mocked mail and fail", func() {
			f, _ := mockSend(errSendingMail)
			mailSender := sender(f)
			recErr := mailSender.Send("hello subject", "hello content")
			So(recErr, ShouldEqual, errSendingMail)
		})
	})
}
