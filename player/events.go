package player

import (
	"bufio"
	"encoding/binary"
	"net"
	"regexp"
	"sunshine/config"
)

// evLogin is called when a player sends a TCPMsgLogin message
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

	// Read Name
	str, err := bufio.NewReader(player.TCPConn).ReadString('\000')
	if err != nil {
		player.Kill(players)
		return
	}

	// Clean up name
	for _, char := range str {
		if name := player.Name + string(char); regexp.MustCompile(config.PlayerNamePattern).MatchString(name) {
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

	// Bitch I'm Verified! https://youtu.be/BpJZAKy3-EI
	player.State = StateVerified
}
