package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const usage = "arguments should be a single port number to listen on\n"

func getPort() int {
	if len(os.Args) != 2 {
		os.Stderr.WriteString(usage)
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port <= 0 || port > 65535 {
		os.Stderr.WriteString(usage)
		os.Exit(1)
	}
	return port
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func main() {
	port := getPort()
	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: port})
	if err != nil {
		fatal(err)
	}
	defer udpConn.Close()

	buffer := make([]byte, 65535)

	for {
		n, err := udpConn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
			continue
		}
		if _, err := os.Stdout.Write(buffer[:n]); err != nil {
			fatal(err)
		}
		if n == 0 || buffer[n-1] != '\n' {
			os.Stdout.Write([]byte{'\n'})
		}
	}
}
