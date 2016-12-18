# Gomail
Gomail is a wrapper around [smtp.SendMail](https://golang.org/pkg/net/smtp/#SendMail).
   
[![Build Status](https://travis-ci.org/uknth/gomail.svg?branch=master)](https://travis-ci.org/uknth/gomail)

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