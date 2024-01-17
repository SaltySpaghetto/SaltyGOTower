package player

import (
	"GOTower/config"
	"encoding/binary"
	"math"
)

// TCP Message Constants
const (
	TCPMsgLogin = 0
	// TCPMsgChat = 1
	// TCPMsgRoomChange = 2
	TCPMsgUUID       = 3
	TCPMsgPlayerLeft = 4
	TCPMsgUdpReady   = 5
)

// Data represents the data constantly updated by a UDP Client.
type Data struct {
	UUID string

	X float32
	Y float32

	Spr     int32
	Frame   uint8
	XScale  int8
	Room    uint16
	Palette uint8
	Name    string
}

// ToBytes converts Data to a byte slice
func (d *Data) ToBytes() []byte {
	b := make([]byte, config.UDPDatagramSize)
	copy(b[0:37], d.UUID)
	binary.LittleEndian.PutUint32(b[37:41], math.Float32bits(d.X))
	binary.LittleEndian.PutUint32(b[41:45], math.Float32bits(d.Y))
	binary.LittleEndian.PutUint32(b[45:49], uint32(d.Spr))
	b[49] = d.Frame
	b[50] = byte(d.XScale)
	binary.LittleEndian.PutUint16(b[51:53], d.Room)
	b[53] = d.Palette
	b = append(b, []byte(d.Name)...)
	return append(b, '\000')
}

// DataFromBytes converts a byte slice to Data
func DataFromBytes(b []byte) Data {
	var data Data
	data.UUID = string(b[0:36])
	data.X = math.Float32frombits(binary.LittleEndian.Uint32(b[37:41]))
	data.Y = math.Float32frombits(binary.LittleEndian.Uint32(b[41:45]))
	data.Spr = int32(binary.LittleEndian.Uint32(b[45:49]))
	data.Frame = b[49]
	data.XScale = int8(b[50])
	data.Room = binary.LittleEndian.Uint16(b[51:53])
	data.Palette = b[53]
	return data
}
