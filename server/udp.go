package server

import (
	"GOTower/config"
	"GOTower/player"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// initializeUDP is the main UDP Component Handler for GoTower.
func (s *Server) initializeUDP() {
	fmt.Printf(config.LangUdpListening, s.UDPPort)
	defer fmt.Print(config.LangUdpClosed)

	// IPv4 *only* because IPv6 is overcomplicating things for this type of application.
	listener, err := net.ListenPacket("udp4", fmt.Sprintf(":%d", s.UDPPort))
	if err != nil {
		fmt.Print(config.LangUdpOpenErr)
		os.Exit(1) // Exit because this isn't the main thread.
	}

	// Deferred so the OS doesn't have to free the socket/port.
	defer func(listener net.PacketConn) {
		_ = listener.Close()
	}(listener)

	for {
		// SEC - Read Packet
		buffer := make([]byte, config.UDPDatagramSize)
		n, addr, err := listener.ReadFrom(buffer)
		if err != nil {
			fmt.Print(config.LangUdpClosed)
			return
		}

		// If n is equal to 2, this is a "hole punch message" used for UDP hole punching.
		// if it is not equal to 2, it MUST be 54, as the only other type of UDP message
		// is the player update message, and it is strictly 54 bytes.
		if n == config.UDPHolepunchSize {
			port := binary.LittleEndian.Uint16(buffer)

			for _, p := range s.Players.Map {
				// Match IP:PORT to the correct player.
				if p.UDPAddr.String() == addr.String() && !p.UDPReady && p.UDPAddr.Port == int(port) {
					b := make([]byte, config.UDPHolepunchSize) //! Response is simply the server's port. May not be necessary?
					binary.LittleEndian.PutUint16(b, uint16(s.UDPPort))
					_, err := listener.WriteTo(b, addr)
					if err != nil {
						p.Kill(s.Players)
						continue
					}
					continue
				}
			}
		} else if n != config.UDPDatagramSize {
			fmt.Printf(config.PrefixError+"UDP Component received a malformed packet from %s.\n", addr.String())
			continue
		}

		// Data is converted from a []byte to Data.
		data := player.DataFromBytes(buffer)

		// SEC - Player Data Update
		s.Players.Mutex.Lock()
		for _, p := range s.Players.Map {
			ip := strings.Split(addr.String(), ":")[0]                    // Proper IP string without port.
			port, _ := strconv.Atoi(strings.Split(addr.String(), ":")[1]) // The port *should* never be anything but.

			// Match player via its IP:PORT combo.
			if p.UDPAddr.IP.String() == ip && p.UDPAddr.Port == port {
				data.Name = p.Name // Name should be static, it can't be changed by the player.

				// Rooms *mustn't* be apart. There was a major error in the previous versions
				// where rooms would be mixed randomly and people wouldn't appear. This is an
				// issue I'm not willing to encounter again.
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
				s.Players.Broadcast(p.Data, listener) // Let the others feast upon the freshly hunted data.
				break
			}
		}
		s.Players.Mutex.Unlock()
	}
}
