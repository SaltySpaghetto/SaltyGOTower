package player

import (
	"encoding/binary"
	"math"
)

// TCP Message Constants
const (
	TCPMsgLogin = iota
	TCPMsgChat
	TCPMsgRoomChange
	TCPMsgUUID
)

// Data updated consistently by UDP packets
type Data struct {
	UUID string

	X float32
	Y float32

	Spr    int32
	Frame  uint8
	XScale uint8
}

// ToBytes converts Data to a byte slice
func (d *Data) ToBytes() []byte {
	b := make([]byte, 51)
	copy(b[0:37], d.UUID)
	binary.LittleEndian.PutUint32(b[37:41], math.Float32bits(d.X))
	binary.LittleEndian.PutUint32(b[41:45], math.Float32bits(d.Y))
	binary.LittleEndian.PutUint32(b[45:49], uint32(d.Spr))
	b[49] = d.Frame
	b[50] = d.XScale
	return b
}

// DataFromBytes converts a byte slice to Data
func DataFromBytes(b []byte) Data {
	var data Data
	data.UUID = string(b[0:36])
	data.X = math.Float32frombits(binary.LittleEndian.Uint32(b[37:41]))
	data.Y = math.Float32frombits(binary.LittleEndian.Uint32(b[41:45]))
	data.Spr = int32(binary.LittleEndian.Uint32(b[45:49]))
	data.Frame = b[49]
	data.XScale = b[50]
	return data
}
