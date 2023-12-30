package main

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2" //calling our 3rd party mail package
)

func listenForMail() {

	go func() {
		for {
			msg := <-app.MailChan //Passing our maildata channel to variable msg
			sendMsg(msg)          //sending our mail using the below function
		}
	}()

}

func sendMsg(m *MailData) {

	//Setting up smtp mail server to email on order posts
	//server login and and security details
	server := mail.NewSMTPClient()
	server.Host = "smtp.mail.yahoo.com"
	server.Port = 587
	server.Username = "wattehgreens@yahoo.com"
	server.Password = "vyxttjvjrfxfauov"
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = true
	server.ConnectTimeout = 40 * time.Second
	server.SendTimeout = 40 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextPlain, string(m.Content))

	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}
	// the part that will finally send the mail
	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent!")
	}

}
