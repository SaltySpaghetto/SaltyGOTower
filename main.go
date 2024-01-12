package main

import (
	"sunshine/config"
	"sunshine/server"
)

// "BOOOOOORING" - Omocat
func main() {
	srv := server.NewServer(config.TcpDefaultPort, config.UdpDefaultPort)
	srv.Listen()
}
