package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)


func main () {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run portscan.go <host> <start> <end>")
		fmt.Println("Example: go run portscan.go example.com 1 10000")
		os.Exit(1)
	}

	host := os.Args[1]
	startPort, _ := strconv.Atoi(os.Args[2])
	endPort, _ := strconv.Atoi(os.Args[3])

	var wg sync.WaitGroup

	fmt.Printf("Scanning %s ports %d %d...\n\n", host, startPort, endPort)

	for port := startPort; port <= endPort; port++ {
		wg.Add(1) // wg counter +1
		go isPortOpen(host, port, &wg)
	}

	wg.Wait() // wg counter = 0
	fmt.Println("\nScan completed!")
}

func isPortOpen(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done() // wg container -1
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err == nil {
		// port open
		conn.Close()
		fmt.Printf("Port %d open\n", port)
	}
}
