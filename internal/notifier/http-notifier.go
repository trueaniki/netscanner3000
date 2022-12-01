package notifier

import (
	"log"
)

type HttpNotifier struct {
	Url    string
	Method string
}

func (n *HttpNotifier) Notify(message string) error {
	log.Println("Sending http request")
	return nil
}
