package server

import (
	"encoding/binary"
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
		buffer := make([]byte, 54)
		n, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			fmt.Print(config.LangUdpClosed)
			return
		}

		// Size must be 53 bytes PERIOD SLAY QUEEN
		if n == 2 {
			port := binary.LittleEndian.Uint16(buffer)
			for _, p := range s.Players.Map {
				if p.UDPAddr.String() == addr.String() && !p.UDPReady && p.UDPAddr.Port == int(port) {
					b := make([]byte, 2)
					binary.LittleEndian.PutUint16(b, s.PortUDP)
					_, err := listener.WriteTo(b, addr)
					if err != nil {
						p.Kill(s.Players)
						continue
					}
					continue
				}
			}
		} else if n != 54 {
			continue
		}

		data := player.DataFromBytes(buffer)

		// Update player data
		s.Players.Mutex.Lock()
		for _, p := range s.Players.Map {
			ip := strings.Split(addr.String(), ":")[0]
			port, _ := strconv.Atoi(strings.Split(addr.String(), ":")[1])
			if p.UDPAddr.IP.String() == ip && p.UDPAddr.Port == port {
				data.Name = p.Name

				if data.Room != p.Data.Room {
					for _, p2 := range s.Players.Map {
						if p2.Data.Room == p.Data.Room {
							b := make([]byte, 38)
							b[0] = player.TCPMsgPlayerLeft
							copy(b[1:37], p.UUID)
							_, err := p2.TCPConn.Write(b)
							if err != nil {
								p2.Kill(s.Players)
								continue
							}
						}
					}
				}

				p.Data = data
				s.Players.Broadcast(p.Data, listener)
				break
			}
		}
		s.Players.Mutex.Unlock()
	}
}
