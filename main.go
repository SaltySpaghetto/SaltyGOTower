package main

import (
	"GOTower/config"
	"GOTower/server"
)

// main is the program's entry point as defined in the Go standard.
func main() {
	srv := server.NewServer(config.TcpDefaultPort, config.UdpDefaultPort)
	srv.Initialize()
}
