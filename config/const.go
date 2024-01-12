package config

// -- Default Conf Constants -- //
const TcpDefaultPort uint16 = 7766
const UdpDefaultPort uint16 = 7767
const PlayerNamePattern = "^[a-zA-Z0-9_]{3,16}$"

// -- Color Constants -- //
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"
const ColorReset = "\033[0m"

// -- Player Language Constants -- //
const PrefixPlayer = ColorCyan + "Player " + ColorWhite + "| " + ColorReset
const LangPlayerJoined = PrefixPlayer + "Player %s joined the game.\n"
const LangPlayerLeft = PrefixPlayer + "Player %s left the game.\n"

// -- Server Language Constants -- //
const Logo = ColorYellow + `                          ,--.    ,--.    
 ,---.,--.,--,--,--, ,---.|  ,---.'--',--,--,,----.
(  .-'|  ||  |      (  .-'|  .-.  ,--|      | .-. :
.-'  ''  ''  |  ||  .-'  '|  | |  |  |  ||  \   --.
'----' '----''--''--'----''--' '--'--'--''--''----'
` + ColorReset

const PrefixLog = ColorCyan + "Log " + ColorWhite + "| " + ColorReset
const PrefixError = ColorRed + "Err " + ColorWhite + "| " + ColorReset

const LangServerWelcome = PrefixLog + "Hello! Welcome to the Sunshine Server!\n"
const LangServerGoodbye = PrefixLog + "Exiting. Goodbye!\n"

const LangTcpListening = PrefixLog + "TCP server listening on TCP:" + ColorYellow + "%d.\n"
const LangTcpClosed = PrefixLog + "TCP server closed.\n"
const LangTcpOpenErr = PrefixError + "TCP server failed to open. Please make sure the port is not in use.\n"
const LangTcpAcceptErr = PrefixError + "TCP server failed to accept a connection.\n"
const LangTcpClientConnected = PrefixLog + "TCP client connected from " + ColorYellow + "%s" + ColorReset + ".\n"

const LangUdpListening = PrefixLog + "UDP server listening on UDP:" + ColorYellow + "%d.\n"
const LangUdpClosed = PrefixLog + "UDP server closed.\n"
const LangUdpOpenErr = PrefixError + "UDP server failed to open. Please make sure the port is not in use.\n"
const LangUdpMessageReceived = PrefixLog + "UDP message received from " + ColorYellow + "%s" + ColorReset + ": " + ColorYellow + "%s" + ColorReset + ".\n"
