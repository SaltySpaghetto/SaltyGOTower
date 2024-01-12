package player

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"sunshine/config"
	"sync"
)

const (
	StateConnected = iota
	StateVerified
	// StatePaused
)

// Player data structure
type Player struct {
	UUID    string
	TCPConn net.Conn
	UDPConn net.Conn
	UDPAddr net.UDPAddr
	State   uint8
	Active  bool

	Name string
	Room uint32
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

// Kill kills the player! DIE YOU FUCKING FAGGOT!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (player *Player) Kill(players Players) {
	player.Active = false
	players.Mutex.Lock()
	delete(players.Map, player.UUID)
	players.Mutex.Unlock()
	fmt.Printf(config.LangPlayerLeft, player.UUID)
}

// Listen for TCP Messages
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
		case TCPMsgLogin:
			player.evLogin(players)
			break

		default:
			player.Kill(players)
			return
		}
	}
}

// Players is simply a wrap around a map. That's a funny rhyme, ain't it?
type Players struct {
	Map   map[string]*Player
	Mutex *sync.Mutex
}

// GetPlayerByUUID does just that!
func (players Players) GetPlayerByUUID(uuid string) *Player {
	players.Mutex.Lock()
	defer players.Mutex.Unlock()
	return players.Map[uuid]
}

// Broadcast sends data to every player in a specified room. Used for updating UDP data.
func (players Players) Broadcast(data Data, room uint32) {
	players.Mutex.Lock()
	defer players.Mutex.Unlock()
	for _, p := range players.Map {
		if p.Active && p.State == StateVerified && p.Room == room && p.UUID != data.UUID {
			_, err := fmt.Fprintf(p.UDPConn, "%s", data.ToBytes())
			if err != nil {
				fmt.Println(err)
				p.Kill(players)
			}
		}
	}
}
