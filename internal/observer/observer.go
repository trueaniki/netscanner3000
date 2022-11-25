package observer

import "github.com/Evencaster/netscanner3000/internal/scanner"

type Observer interface {
	Receive(scannerID scanner.ScannerID, data scanner.PortsData)
}
