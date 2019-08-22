package main

import (
	"net"
	"os"
	"strconv"
)

const newline = byte('\n')

var newlineBytes = []byte{byte('\n')}

func getPort() int {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("arguments should be a single port number to listen on\n")
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Stderr.WriteString("arguments should be a single port number to listen on\n")
		os.Exit(1)
	}
	return port
}

func main() {
	port := getPort()
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{Port: port})
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.Write(newlineBytes)
		os.Exit(1)
	}
	defer udpConn.Close()

	buffer := make([]byte, 65535)

	for {
		n, err := udpConn.Read(buffer)
		if err != nil {
			os.Stderr.WriteString("Read error: ")
			os.Stderr.WriteString(err.Error())
			os.Stderr.Write(newlineBytes)
			continue
		}
		os.Stdout.Write(buffer[:n])
		if buffer[n-1] != newline {
			os.Stdout.Write(newlineBytes)
		}
	}
}
