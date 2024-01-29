package player

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"

	"GOTower/constants"
)

// evLogin is an event fired upon receiving a TCPMsgLogin typed message from a Player.
func (player *Player) evLogin(players Players) {
	// You can't log in twice
	if player.State != StateConnected {
		return
	}

	// Read Port Value
	port := make([]byte, 2)
	_, err := player.TCPConn.Read(port)
	if err != nil {
		player.Kill(players)
		return
	}
	player.UDPAddr.Port = int(binary.LittleEndian.Uint16(port))

	// Init UDPConn
	player.UDPConn, err = net.Dial("udp", player.UDPAddr.String())
	if err != nil {
		player.Kill(players)
		return
	}

	// Read Version
	reader := bufio.NewReader(player.TCPConn)
	version, err := reader.ReadString('\000')
	if err != nil {
		player.Kill(players)
		return
	}

	version = version[:len(version)-1]

	// Check Version
	if version != constants.ServerVersion {
		player.TCPConn.Write([]byte{TCPWrongVersion})
		player.Kill(players)
		return
	}

	// Read Name
	str, err := reader.ReadString('\000')
	if err != nil {
		player.Kill(players)
		return
	}

	// Clean up name
	for _, char := range str {
		if name := player.Name + string(char); regexp.MustCompile(constants.PlayerNamePattern).MatchString(name) {
			player.Name = name
		}
	}

	if len(player.Name) < 1 {
		player.Kill(players)
		return
	}

	// UUID is 36 characters long, and the null terminator is 1 character long.
	// Factor in the UUID message type byte, and you get 38. Keep that in mind.
	b := make([]byte, 38)
	b[0] = TCPMsgUUID
	copy(b[1:37], player.UUID)
	b[37] = '\000'
	_, err = player.TCPConn.Write(b)
	if err != nil {
		player.Kill(players)
		return
	}

	player.State = StateVerified

	// Alert everyone of the joining player.

	players.Mutex.Lock()
	players.BroadcastTCP(ChatMessage{
		Msg:  fmt.Sprintf("Player %s entered the tower.", player.Name),
		Name: "Server",
	}.ToBytes())
	players.Mutex.Unlock()
}

// evChat is fired when a Player sends a chat message. It cleans the message and broadcasts it to everyone else.
func (player *Player) evChat(players Players) {
	reader := bufio.NewReader(player.TCPConn)
	str, err := reader.ReadString('\000')
	if err != nil {
		player.Kill(players)
		return
	}

	// Clean the string
	for _, char := range str {
		if newstr := str + string(char); regexp.MustCompile(constants.PlayerChatPattern).MatchString(str) {
			str = newstr
		}
	}

	msg := ChatMessage{
		Msg:  str,
		Name: player.Name,
	}

	fmt.Printf("%s (%s): %s\n", msg.Name, player.UUID, msg.Msg)

	players.Mutex.Lock()
	players.BroadcastTCP(msg.ToBytes())
	players.Mutex.Unlock()
}
