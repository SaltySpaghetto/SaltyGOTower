package config

import _ "embed"

// -- Default Conf Constants -- //
const ServerVersion string = "24w04a"
const TcpDefaultPort uint16 = 7766
const UdpDefaultPort uint16 = 7767
const PlayerNamePattern = "^[a-zA-Z0-9_]{1,16}$"

// SEC - Message Sizes
const UDPDatagramSize = 59
const UDPHolepunchSize = 2

// SEC - Colors
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"
const ColorReset = "\033[0m"

// SEC - Player Language
const PrefixPlayer = ColorCyan + "Player " + ColorWhite + "| " + ColorReset
const LangPlayerLeft = PrefixPlayer + "Player " + ColorYellow + "%s left the game.\n" + ColorReset

// SEC - Server Language
//
//go:embed logo.txt
var Logo string

const Name = ColorCyan + "≡GO" + ColorPurple + "Tower" + ColorReset

const PrefixLog = ColorCyan + "Log " + ColorPurple + "| " + ColorReset
const PrefixError = ColorRed + "Err " + ColorPurple + "| " + ColorReset

const LangServerWelcome = PrefixLog + "Initializing " + Name + ".\n"
const LangServerGoodbye = PrefixLog + "Exiting now.\n"

const LangTcpListening = PrefixLog + "TCP Component initialized on TCP:" + ColorYellow + "%d.\n"
const LangTcpClosed = PrefixLog + "TCP Component exited.\n"
const LangTcpOpenErr = PrefixError + "TCP Component failed to start. Please ensure the port specified is not in use.\n"
const LangTcpAcceptErr = PrefixError + "TCP Component failed to accept a connection.\n"
const LangTcpClientConnected = PrefixLog + "TCP Component accepted a connection from " + ColorYellow + "%s" + ColorReset + ".\n"

const LangUdpListening = PrefixLog + "UDP Component initialized on UDP:" + ColorYellow + "%d.\n"
const LangUdpClosed = PrefixLog + "UDP Component exited.\n"
const LangUdpOpenErr = PrefixError + "UDP Component failed to start. Please ensure the port is not in use.\n"

const LangConfigNotfoundErr = PrefixError + "Could not find the config file.\n"
const LangConfigFormatErr = PrefixError + "Config file is not formatted correctly.\n"
