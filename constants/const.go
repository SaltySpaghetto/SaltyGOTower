package constants

import _ "embed"

// SEC - Server Data
const ServerVersion string = "24w04c"

// SEC - Default Server Settings
const TcpDefaultPort uint16 = 7766
const UdpDefaultPort uint16 = 7767
const PlayerNamePattern = "^[a-zA-Z0-9_]{1,16}$"
const PlayerChatPattern = "^[a-zA-Z0-9_ .><?()!/]{1,64}$"
