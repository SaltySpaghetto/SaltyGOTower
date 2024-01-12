package server

import (
	"fmt"
	"sunshine/config"
	"sunshine/player"
	"sync"
)

// Server data structure
type Server struct {
	PortTCP uint16
	PortUDP uint16

	MaxPlayers uint16
	Players    player.Players
}

// NewServer creates a new server object
func NewServer(portTCP uint16, portUDP uint16) *Server {
	return &Server{
		PortTCP: portTCP,
		PortUDP: portUDP,

		MaxPlayers: 255,
		Players: player.Players{
			Map:   make(map[string]*player.Player),
			Mutex: &sync.Mutex{},
		},
	}
}

// Listen for incoming TCP connections
func (s *Server) Listen() {
	fmt.Print(config.Logo)
	fmt.Print(config.LangServerWelcome)
	defer fmt.Print(config.LangServerGoodbye)

	go s.listenUDP()
	s.listenTCP()
}
