package player

import (
	"GOTower/config"
	"fmt"
	"github.com/google/uuid"
	"net"
	"sync"
)

const (
	StateConnected = iota
	StateVerified
)

// Player data structure
type Player struct {
	UUID     string
	TCPConn  net.Conn
	UDPConn  net.Conn
	UDPReady bool
	UDPAddr  net.UDPAddr
	State    uint8
	Active   bool

	Name string
	Data Data
}

// NewPlayer creates a new player object
func NewPlayer(conn net.Conn) *Player {
	return &Player{
		UUID:    uuid.New().String(),
		TCPConn: conn,
		UDPConn: nil,
		UDPAddr: net.UDPAddr{IP: conn.RemoteAddr().(*net.TCPAddr).IP},
		State:   StateConnected,
		Active:  true,

		Name: "",
		Data: Data{},
	}
}

// Kill is an eloquently named function that disconnects the player
// and alerts all players in the same room of their sudden absence.
func (player *Player) Kill(players Players) {
	player.Active = false
	players.Mutex.Lock()

	// 38 = Type + UUID + \000.
	b := make([]byte, 38)
	b[0] = TCPMsgPlayerLeft
	copy(b[1:37], player.UUID)

	for _, p := range players.Map {
		// Alert every player in the same room.
		if p.Data.Room == player.Data.Room {
			_, _ = p.TCPConn.Write(b)
		}
	}
	delete(players.Map, player.UUID)
	players.Mutex.Unlock()
	fmt.Printf(config.LangPlayerLeft, player.UUID)
}

// Listen is how a Player thread listens for incoming TCP messages.
func (player *Player) Listen(players Players) {
	defer func() {
		_ = player.TCPConn.Close()
	}()

	for player.Active {
		msgType := make([]byte, 1)

		_, err := player.TCPConn.Read(msgType)
		if err != nil {
			player.Kill(players)
			fmt.Println(err)
			return
		}

		switch msgType[0] {
		case TCPMsgLogin: // Fired on Player Login ONLY.
			player.evLogin(players)
			break

		case TCPMsgUdpReady: // Fired when UDP hole punching succeeds.
			player.UDPReady = true
			break

		default:
			player.Kill(players)
			return
		}
	}
}

// Players is simply a wrap around a map. Quite the rhyme, isn't it?
type Players struct {
	Map   map[string]*Player
	Mutex *sync.Mutex
}

// Broadcast sends data to every player in a specified room. Used for updating UDP data. MUST lock the mutex before calling this function.
func (players Players) Broadcast(data Data, listener net.PacketConn) {
	for _, p := range players.Map {
		if p.Active && p.State == StateVerified && p.Data.Room == data.Room && p.UUID != data.UUID {
			_, err := listener.WriteTo(data.ToBytes(), &p.UDPAddr)
			if err != nil {
				fmt.Println(err)
				p.Kill(players)
			}
		}
	}
}
