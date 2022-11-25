package main

import (
	"os"
	"time"

	"github.com/Evencaster/netscanner3000/pkg/config"
	"github.com/Evencaster/netscanner3000/pkg/notifier"
	"github.com/Evencaster/netscanner3000/pkg/observer"
	"github.com/Evencaster/netscanner3000/pkg/scanner"
)

func main() {
	c := config.Parse(os.Args[1])

	o, err := initObserver(c)
	if err != nil {
		panic("Cannot initialze daemon")
	}
	scanners := initScanners(c)

	n := &notifier.EmailNotifier{
		From:     c.Notifiers.Email.From,
		To:       []string{c.Notifiers.Email.To},
		Password: c.Notifiers.Email.Password,
		SmtpHost: c.Notifiers.Email.SmtpHost,
		SmtpPort: c.Notifiers.Email.SmtpPort,
	}

	o.Scanners = scanners
	o.Listeners = []notifier.Notifier{n}

	o.Start()
}

func initObserver(c config.Config) (*observer.Observer, error) {
	interval, err := time.ParseDuration(c.Daemon.Interval)
	if err != nil {
		return nil, err
	}
	d := &observer.Observer{
		Interval: interval,
	}
	return d, nil
}

func initScanners(c config.Config) []*scanner.Scanner {
	scanners := make([]*scanner.Scanner, len(c.Scanners))
	for i, s := range c.Scanners {
		scanners[i] = &scanner.Scanner{
			Adress:    s.Host,
			StartPort: s.StartPort,
			EndPort:   s.EndPort,
			Network:   s.Network,
		}
	}
	return scanners
}
