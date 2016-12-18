# Gomail
Gomail is a wrapper around [smtp.SendMail](https://golang.org/pkg/net/smtp/#SendMail).
   
[![Build Status](https://travis-ci.org/uknth/gomail.svg?branch=master)](https://travis-ci.org/uknth/gomail)
[![GoDoc](https://godoc.org/github.com/tj/go-rle?status.svg)](https://godoc.org/github.com/tj/go-rle)
[![](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

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

---
> [uknth.me](http://uknth.me) &nbsp;&middot;&nbsp;
> GitHub [@uknth](https://github.com/uknth) &nbsp;&middot;&nbsp;
> Twitter [@uknth](https://twitter.com/uknth)