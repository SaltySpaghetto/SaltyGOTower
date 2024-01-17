package server

import (
	"GOTower/config"
	"GOTower/player"
	"fmt"
	"sync"
)

// Server is a structure that stores important data about the current server instance.
type Server struct {
	TCPPort uint
	UDPPort uint

	MaxPlayers uint
	Players    player.Players
}

// NewServer initializes a server instance using a config.ini file.
func NewServer(portTCP uint16, portUDP uint16) *Server {
	conf := config.LoadConfig("config.ini")
	return &Server{
		TCPPort: conf.TCPPort,
		UDPPort: conf.UDPPort,

		MaxPlayers: conf.MaxPlayers,
		Players: player.Players{
			Map:   make(map[string]*player.Player),
			Mutex: &sync.Mutex{},
		},
	}
}

// Initialize initializes both the TCP and UDP Components
func (s *Server) Initialize() {
	fmt.Print(config.Logo)
	fmt.Print(config.LangServerWelcome)
	defer fmt.Print(config.LangServerGoodbye)

	go s.initializeUDP()
	s.initializeTCP()
}
