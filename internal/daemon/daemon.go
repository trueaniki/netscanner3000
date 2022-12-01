package daemon

import (
	"github.com/Evencaster/netscanner3000/internal/observer"
	"github.com/Evencaster/netscanner3000/internal/scanner"
)

type Startable interface {
	Start()
}

type Daemon struct {
	Scanners  []scanner.Runnable
	Observers []observer.Observer
}

func (d *Daemon) Start() {
	panic("Not implemented")
}

func (d *Daemon) tick() {
	for _, s := range d.Scanners {
		res := s.Run()
		d.emit(s.ID, res)
	}
}

func (d *Daemon) emit(scannerID scanner.ScannerID, data scanner.PortsData) {
	for _, o := range d.Observers {
		o.Receive(scannerID, data)
	}
}
