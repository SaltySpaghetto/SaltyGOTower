package server

import (
	"fmt"
	"net"
	"os"
	"sunshine/config"
	"sunshine/player"
)

func (s *Server) listenTCP() {
	fmt.Printf(config.LangTcpListening, s.PortTCP)
	defer fmt.Print(config.LangTcpClosed)

	// I looooooove ipv4!
	listener, err := net.Listen("tcp4", fmt.Sprintf(":%d", s.PortTCP))

	if err != nil {
		fmt.Print(config.LangTcpOpenErr)
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	for {
		// YEEHAW
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(config.LangTcpAcceptErr)
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf(config.LangTcpClientConnected, conn.RemoteAddr().String())

		p := player.NewPlayer(conn)
		s.Players.Mutex.Lock()
		s.Players.Map[p.UUID] = p
		s.Players.Mutex.Unlock()
		go p.Listen(s.Players)
	}
}
