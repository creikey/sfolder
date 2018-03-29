package main

import (
	"bufio"
	"net"
)

const (
	port string = ":3333"
)

// Client is a standard client
type Client struct {
	IP         string
	Connection net.Conn
	RWriter    *bufio.ReadWriter
}

// CreateClient connects to an ip and returns a client
func CreateClient(ip string) (Client, error) {
	conn, err := net.Dial("tcp", ip+port)
	if err != nil {
		return Client{}, err
	}
	return Client{ip, conn, bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))}, nil
}
