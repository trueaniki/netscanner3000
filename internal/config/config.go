package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
	"github.com/Evencaster/netscanner3000/internal/daemon"
	"github.com/Evencaster/netscanner3000/internal/scanner"
)

func Parse(filepath string) (c Config) {
	configData, err := os.ReadFile(filepath)
	if err != nil {
		panic("Cannot read config file")
	}

	err = yaml.Unmarshal(configData, &c)
	if err != nil {
		panic("Cannot parse config file")
	}
	return
}

type DaemonConfig struct {
	Interval string `yaml:"interval"`
}

type ScannerConfig struct {
	Host      string `yaml:"host"`
	StartPort int    `yaml:"startPort"`
	EndPort   int    `yaml:"endPort"`
	Network   string `yaml:"network"`
}

type ObserverConfig struct {
	Notify struct {
		Smtp SmtpConfig `yaml:"smtp"`
	} `yaml:"notify"`
}

type SmtpConfig struct {
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	SmtpHost string `yaml:"smtpHost"`
	SmtpPort string `yaml:"smtpPort"`
	To       string `yaml:"to"`
}

type Config struct {
	Daemon    DaemonConfig    `yaml:"daemon"`
	Scanners  []ScannerConfig `yaml:"scanners"`
	Observers struct {
		OnChange []ObserverConfig `yaml:"onChange"`
	} `yaml:"observers"`
}

func (c *Config) GetDaemon() *daemon.Daemon {
	interval, err := time.ParseDuration(c.Daemon.Interval)
	if err != nil {
		panic("Cannot initialize daemon")
	}
	return &daemon.Daemon{
		Interval: interval,
	}
}

func (c *Config) GetScanners() []*scanner.Scanner {

}

func (c *Config) GetObservers() {

}
