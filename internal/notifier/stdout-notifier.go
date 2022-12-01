package notifier

import "log"

type StdoutNotifier struct{}

func (n *StdoutNotifier) Notify(message string) error {
	log.Println(message)
	return nil
}
