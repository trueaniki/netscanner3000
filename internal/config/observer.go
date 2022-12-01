package config

import "github.com/Evencaster/netscanner3000/internal/observer"

type ObserverConfig struct {
	Notify NotifierConfig `yaml:"notify"`
}

func (c *Config) GetObservers() []observer.Observer {
	observers := make([]observer.Observer, 0, len(c.Observers.OnChange))

	for _, o := range c.Observers.OnChange {
		observers = append(observers, &observer.ChangeObserver{
			Notifyier: getNotifier(o.Notify),
		})
	}

	return observers
}
