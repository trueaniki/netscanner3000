package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type ScannerID = int
type PortsData = map[int]bool

type Scanner struct {
	ID        ScannerID
	Adress    string
	StartPort int
	EndPort   int
	Network   string
	Timeout   time.Duration
}

func (s *Scanner) connect(port int) bool {
	addr := fmt.Sprintf("%s:%d", s.Adress, port)
	_, err := net.DialTimeout(s.Network, addr, 5*time.Second)
	return err == nil
}

func (s *Scanner) Run() (res PortsData) {
	m := sync.Mutex{}
	res = make(PortsData)

	wg := sync.WaitGroup{}
	wg.Add(s.EndPort - s.StartPort + 1)
	for p := s.StartPort; p <= s.EndPort; p++ {
		go func(p int) {
			isOpen := s.connect(p)
			m.Lock()
			res[p] = isOpen
			m.Unlock()
			wg.Done()
		}(p)
	}
	wg.Wait()
	return
}
