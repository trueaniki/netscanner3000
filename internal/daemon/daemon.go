package daemon

import (
	"time"

	"github.com/Evencaster/netscanner3000/internal/observer"
	"github.com/Evencaster/netscanner3000/internal/scanner"
)

type Daemon struct {
	Interval  time.Duration
	Scanners  []*scanner.Scanner
	Observers []observer.Observer
}

func (d *Daemon) Start() {
	for _, s := range d.Scanners {
		res := s.Run()
		d.emit(s.ID, res)
	}

	time.Sleep(d.Interval)
}

func (d *Daemon) emit(scannerID scanner.ScannerID, data scanner.PortsData) {
	for _, o := range d.Observers {
		o.Receive(scannerID, data)
	}
}
