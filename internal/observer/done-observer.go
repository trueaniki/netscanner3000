package observer

import (
	"fmt"
	"strings"

	"github.com/Evencaster/netscanner3000/internal/notifier"
	"github.com/Evencaster/netscanner3000/internal/scanner"
)

type DoneObserver struct {
	Notifyier notifier.Notifier
}

func (o *DoneObserver) Receive(scannerID scanner.ScannerID, data scanner.PortsData) {
	// TODO: optimize
	messages := make([]string, 0)

	for port, state := range data {
		messages = append(
			messages,
			fmt.Sprintf("Port %d is %t", port, state),
		)
	}
	o.Notifyier.Notify(strings.Join(messages, "\n"))
}
