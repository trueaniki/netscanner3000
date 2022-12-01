package config

import (
	"time"

	"github.com/Evencaster/netscanner3000/internal/daemon"
	"github.com/go-co-op/gocron"
)

type DaemonConfig struct {
	Interval string `yaml:"interval"`
	Cron     string `yaml:"cron"`
}

func (c *Config) GetDaemon() daemon.Startable {
	if c.Daemon.Interval != "" {
		interval, err := time.ParseDuration(c.Daemon.Interval)
		if err != nil {
			panic("Cannot initialize daemon")
		}
		return &daemon.IntervalDaemon{
			Interval: interval,
			Daemon: daemon.Daemon{
				Scanners:  c.GetScanners(),
				Observers: c.GetObservers(),
			},
		}
	}
	if c.Daemon.Cron != "" {
		s := gocron.NewScheduler(time.UTC).Cron(c.Daemon.Cron)
		return &daemon.CronDaemon{
			Scheduler: s,
			Daemon: daemon.Daemon{
				Scanners:  c.GetScanners(),
				Observers: c.GetObservers(),
			},
		}
	}
	return nil
}
