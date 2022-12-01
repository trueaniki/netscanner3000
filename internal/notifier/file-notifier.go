package notifier

import "os"

type FileNotifier struct {
	Output string
}

func (n *FileNotifier) Notify(message string) error {
	f, err := os.OpenFile(n.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(message); err != nil {
		return err
	}
	return nil
}
