package daemon

import "time"

type IntervalDaemon struct {
	Daemon
	Interval time.Duration
}

func (d *IntervalDaemon) Start() {
	d.tick()
	time.Sleep(d.Interval)
}
