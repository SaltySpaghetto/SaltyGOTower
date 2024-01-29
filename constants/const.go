package constants

import _ "embed"

// SEC - Server Data
const ServerVersion string = "24w04d"

// SEC - Default Server Settings
const (
	TcpDefaultPort    uint16 = 7766
	UdpDefaultPort    uint16 = 7767
	PlayerNamePattern        = "^[a-zA-Z0-9_]{1,16}$"
	PlayerChatPattern        = "^[a-zA-Z0-9_ .><?()!/]{1,64}$"
)
