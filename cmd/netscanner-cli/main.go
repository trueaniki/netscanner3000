package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Evencaster/netscanner3000/internal/scanner"
)

func main() {
	startPort, _ := strconv.Atoi(os.Args[2])
	endPort, _ := strconv.Atoi(os.Args[3])

	s := scanner.Scanner{
		Adress:    os.Args[1],
		StartPort: int(startPort),
		EndPort:   int(endPort),
		Protocol:  os.Args[4],
	}
	fmt.Println(s.Run())
}
