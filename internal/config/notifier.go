package config

import "github.com/Evencaster/netscanner3000/internal/notifier"

type NotifierConfig struct {
	Stdout *StdoutNotifierConfig `yaml:"stdout"`
	File   *FileNotifierConfig   `yaml:"file"`
	Smtp   *SmtpNotifierConfig   `yaml:"smtp"`
	Http   *HttpNotifierConfig   `yaml:"http"`
}

type StdoutNotifierConfig struct{}

type FileNotifierConfig struct {
	Output string `yaml:"output"`
}

type SmtpNotifierConfig struct {
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	SmtpHost string `yaml:"smtpHost"`
	SmtpPort string `yaml:"smtpPort"`
	To       string `yaml:"to"`
}

type HttpNotifierConfig struct {
	Url    string `yaml:"url"`
	Method string `yaml:"method"`
}

func getNotifier(c NotifierConfig) notifier.Notifier {
	//  Stdout notify
	if c.Stdout != nil {
		return &notifier.StdoutNotifier{}
	}

	// File notify
	if c.File != nil {
		return &notifier.FileNotifier{
			Output: c.File.Output,
		}
	}

	// Smtp notify
	if c.Smtp != nil {
		return &notifier.SmtpNotifier{
			From:     c.Smtp.From,
			To:       c.Smtp.To,
			Password: c.Smtp.Password,
			SmtpHost: c.Smtp.SmtpHost,
			SmtpPort: c.Smtp.SmtpPort,
		}
	}

	// Http notify
	if c.Http != nil {
		return &notifier.HttpNotifier{
			Url:    c.Http.Url,
			Method: c.Http.Method,
		}
	}
	return nil
}
