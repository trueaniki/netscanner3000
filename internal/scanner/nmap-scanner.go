package scanner

import (
	"log"
	"strconv"

	"github.com/Ullaakut/nmap"
)

type NmapScanner struct {
	Scanner
}

func (s *NmapScanner) Run() {
	ports := ""

	for i := s.StartPort; i <= s.EndPort; i++ {
		ports = ports + strconv.Itoa(i)
		if i != s.EndPort {
			ports = ports + ","
		}
	}

	nmapScanner, err := nmap.NewScanner(
		nmap.WithTargets(s.Adress),
		nmap.WithPorts(ports),
	)
	if err != nil {
		log.Fatal(err)
	}

	nmapScanner.Run()
}
