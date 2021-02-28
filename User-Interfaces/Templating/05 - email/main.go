package main

import (
	"bytes"
	"net/smtp"
	"strconv"
	"text/template"
)

type EmailMessage struct {
	From, Subject, Body string
	To                  []string
}
type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
{{.Body}}`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	message := &EmailMessage{
		From:    "peopiayo@peopyao.pi",
		To:      []string{"peopiayothesecond@peopyao.pi"},
		Subject: "Ping-Poke",
		Body:    "Where is the money, Lamborginsky?",
	}
	var buffy bytes.Buffer

	t.Execute(&buffy, message)

	authCreds := &EmailCredentials{
		Username: "Usernamedamad",
		Password: "NotInTheSudoersFile",
		Server:   "smtp.domain.zn",
		Port:     25,
	}
	auth := smtp.PlainAuth("", authCreds.Username, authCreds.Password, authCreds.Server)
	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port), auth, message.From, message.To, buffy.Bytes())
}
