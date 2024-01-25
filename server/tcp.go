package server

import (
	"GOTower/constants"
	"GOTower/player"
	"fmt"
	"net"
)

// initializeTCP is the main TCP Component Handler for GoTower.
func (s *Server) initializeTCP() {
	fmt.Printf(constants.LangTcpListening, s.TCPPort)
	defer fmt.Print(constants.LangTcpClosed)

	// IPv4 *only* because IPv6 is overcomplicating things for this type of application.
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", s.TCPPort))
	if err != nil {
		fmt.Print(constants.LangTcpOpenErr)
		fmt.Println(err) // Do not exit because this is the main thread.
		return
	}

	// Deferred so the OS doesn't have to free the socket/port.
	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	for {
		// Accept TCP Connections.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(constants.LangTcpAcceptErr)
			fmt.Println(err)
			return
		}
		fmt.Printf(constants.LangTcpClientConnected, conn.RemoteAddr().String())

		// Initialize the Player who has connected.
		p := player.NewPlayer(conn)
		s.Players.Mutex.Lock()
		s.Players.Map[p.UUID] = p
		s.Players.Mutex.Unlock()
		go p.Listen(s.Players)
	}
}
