package player

import (
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
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

	// Read Name Value (Max 16)
	for {
		// Read 1 Character
		char := make([]byte, 1)
		_, err := player.TCPConn.Read(char)
		if err != nil {
			player.Kill(players)
			fmt.Println(err)
			return
		}

		// Break if character is null
		if char[0] == '\000' {
			break
		}

		// Test if character matches regex
		newName := player.Name + string(char)
		matched, err := regexp.MatchString(`^[a-zA-Z0-9_]+$`, newName)
		if err != nil {
			player.Kill(players)
			fmt.Println(err)
			return
		}

		// Push changes to name value if it matches the regex
		if matched {
			player.Name = newName
		}
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
