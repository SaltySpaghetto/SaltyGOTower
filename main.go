package main

import (
	"GOTower/constants"
	"GOTower/server"
)

// main is the program's entry point as defined in the Go standard.
func main() {
	srv := server.NewServer(constants.TcpDefaultPort, constants.UdpDefaultPort)
	srv.Initialize()
}
