# Gomail
Gomail is a wrapper around [smpt.SendMail](https://golang.org/pkg/net/smtp/#SendMail).

### Usage

```
var mailer = gomail.New(&gomail.Config{
	UserName: "xyz@abc.com",
	Password: "password",
	ServerHost: "host",
	ServerPort: "port",
	SenderAddr: "xyz@abc.com"
	Prefix: "Some Subject Line Prefix - "
	Receivers: []string {  
		"123@abc.com",
		"456@abc.com"
	}
})

mailer.Send("Test Subject", "Test Content")

```