package observer

import (
	"fmt"
	"log"
	"strings"

	"github.com/Evencaster/netscanner3000/internal/notifier"
	"github.com/Evencaster/netscanner3000/internal/scanner"
)

type ChangeObserver struct {
	Notifyier notifier.Notifier
	state     map[scanner.ScannerID]scanner.PortsData
}

func (o *ChangeObserver) Receive(scannerID scanner.ScannerID, data scanner.PortsData) {
	if o.state == nil {
		o.state = make(map[scanner.ScannerID]scanner.PortsData)
	}
	prevState := o.state[scannerID]
	newState := data
	changes := []string{}

	for port, prevPortState := range prevState {
		if newState[port] != prevPortState {
			changes = append(
				changes,
				fmt.Sprintf("Port %d was %t, become %t \n", port, prevPortState, newState[port]),
			)
		}
	}
	if o.state[scannerID] == nil {
		o.state[scannerID] = make(scanner.PortsData)
	}
	o.state[scannerID] = newState

	if len(changes) != 0 {
		if err := o.Notifyier.Notify(strings.Join(changes, "\n")); err != nil {
			log.Fatal(err)
		}
	}
}
