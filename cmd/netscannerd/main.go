package main

import (
	"os"

	"github.com/Evencaster/netscanner3000/internal/config"
)

func main() {
	c := config.Parse(os.Args[1])
	d := c.GetDaemon()

	d.Start()
}
