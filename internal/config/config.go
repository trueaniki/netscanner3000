package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func Parse(filepath string) (c Config) {
	configData, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Cannot read config file")
	}

	err = yaml.Unmarshal(configData, &c)
	if err != nil {
		log.Fatal("Cannot parse config file")
	}
	return
}

type Config struct {
	Daemon    DaemonConfig    `yaml:"daemon"`
	Scanners  []ScannerConfig `yaml:"scanners"`
	Observers struct {
		OnChange []ObserverConfig `yaml:"onChange"`
	} `yaml:"observers"`
}
