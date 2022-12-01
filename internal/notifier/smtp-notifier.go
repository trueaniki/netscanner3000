package notifier

import (
	"log"
	"net/smtp"
)

type SmtpNotifier struct {
	From     string
	Password string
	To       string
	SmtpHost string
	SmtpPort string
}

func (n *SmtpNotifier) Notify(message string) error {
	log.Println("Sending email")
	auth := smtp.PlainAuth("", n.From, n.Password, n.SmtpHost)
	err := smtp.SendMail(n.SmtpHost+":"+n.SmtpPort, auth, n.From, []string{n.To}, []byte(message))
	return err
}
