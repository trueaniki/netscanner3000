package config

import (
	"encoding/binary"
	"log"
	"net"
	"time"

	"github.com/Evencaster/netscanner3000/internal/scanner"
	"github.com/google/uuid"
)

type ScannerConfig struct {
	Host      string `yaml:"host"`
	Cidr      string `yaml:"cidr"`
	StartPort int    `yaml:"startPort"`
	EndPort   int    `yaml:"endPort"`
	Protocol  string `yaml:"protocol"`
	Timeout   string `yaml:"timeout"`
}

func (c *Config) GetScanners() []*scanner.NmapScanner {
	scanners := make([]*scanner.NmapScanner, 0, len(c.Scanners))
	for _, s := range c.Scanners {
		timeout, err := time.ParseDuration(s.Timeout)
		if err != nil {
			log.Fatal("Cannot parse timeout string")
		}
		if s.Cidr != "" {
			for _, h := range getAllHostsFromCidr(s.Cidr) {
				scanners = append(scanners, &scanner.NmapScanner{
					Scanner: scanner.Scanner{
						Adress:    h,
						StartPort: s.StartPort,
						EndPort:   s.EndPort,
						Protocol:  s.Protocol,
						Timeout:   timeout,
					},
				})
			}
		}
		scanners = append(scanners, &scanner.NmapScanner{
			Scanner: scanner.Scanner{
				ID:        int(uuid.New().ID()),
				Adress:    s.Host,
				StartPort: s.StartPort,
				EndPort:   s.EndPort,
				Protocol:  s.Protocol,
				Timeout:   timeout,
			},
		})
	}

	return scanners
}

func getAllHostsFromCidr(cidr string) []string {
	_, ipv4Net, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal(err)
	}

	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)

	finish := (start & mask) | (mask ^ 0xffffffff)

	hosts := make([]string, 0, finish-start+1)
	for i := start; i <= finish; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}
	return hosts
}
