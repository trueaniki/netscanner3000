package observer

import (
	"fmt"
	"strings"

	"github.com/Evencaster/netscanner3000/internal/scanner"
)

type ChangeObserver struct {
	state map[scanner.ScannerID]scanner.PortsData
}

func (o *ChangeObserver) Receive(scannerID scanner.ScannerID, data scanner.PortsData) {
	prevState := o.state[scannerID]
	newState := data
	changes := []string{}

	for port, prevPortState := range prevState {
		if newState[port] != prevPortState {
			changes = append(
				changes,
				fmt.Sprintf("Port %d was %t, become %t", port, prevPortState, newState[port]),
			)
		}
	}
	o.state[scannerID] = newState

	if len(changes) != 0 {
		o.notify(strings.Join(changes, "\n"))
	}
}

func (o *ChangeObserver) notify(m string) {

}
