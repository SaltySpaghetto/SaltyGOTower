package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sunshine/config"
	"sunshine/player"
)

// listenUDP listens for UDP packets on the server's UDP port.
func (s *Server) listenUDP() {
	fmt.Printf(config.LangUdpListening, s.PortUDP)
	defer fmt.Print(config.LangUdpClosed)

	// I looooooove ipv4!
	listener, err := net.ListenPacket("udp4", fmt.Sprintf(":%d", s.PortUDP))

	if err != nil {
		fmt.Print(config.LangUdpOpenErr)
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(listener net.PacketConn) {
		_ = listener.Close()
	}(listener)

	for {
		// Read UDP packet
		buffer := make([]byte, 51)
		n, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			fmt.Print(config.LangUdpClosed)
			return
		}

		// Size must be 51 bytes PERIOD SLAY QUEEN
		if n != 51 {
			fmt.Printf(config.LangUdpMessageReceived, addr.String(), strings.Split(string(buffer), "\000")[0])
			continue
		}

		data := player.DataFromBytes(buffer)

		// Update player data
		s.Players.Mutex.Lock()
		for _, p := range s.Players.Map {
			ip := strings.Split(addr.String(), ":")[0]
			port, _ := strconv.Atoi(strings.Split(addr.String(), ":")[1])
			if p.UDPAddr.IP.String() == ip && p.UDPAddr.Port == port {
				p.Data = data
				s.Players.Broadcast(p.Data, p.Room)
				break
			}
		}
		s.Players.Mutex.Unlock()
	}
}
