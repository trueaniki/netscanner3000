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
	Protocol  string
	Timeout   time.Duration
}

type Runnable interface {
	Run() PortsData
}

func (s *Scanner) connect(port int) bool {
	addr := fmt.Sprintf("%s:%d", s.Adress, port)
	_, err := net.DialTimeout(s.Protocol, addr, s.Timeout)
	return err == nil
}

func (s *Scanner) Run() (res PortsData) {
	return s.runSync()
}

func (s *Scanner) runAsync() (res PortsData) {
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

func (s *Scanner) runSync() (res PortsData) {
	res = make(PortsData)
	for p := s.StartPort; p <= s.EndPort; p++ {
		isOpen := s.connect(p)
		res[p] = isOpen
	}

	return
}
