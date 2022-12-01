package daemon

import "github.com/go-co-op/gocron"

type CronDaemon struct {
	Daemon
	Scheduler *gocron.Scheduler
}

func (d *CronDaemon) Start() {
	d.Scheduler.Do(d.tick)
	d.Scheduler.StartBlocking()
}
