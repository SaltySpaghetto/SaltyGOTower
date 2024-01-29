package constants

import _ "embed"

// SEC - Player Language
const (
	PrefixPlayer   = ColorCyan + "Player " + ColorWhite + "| " + ColorReset
	LangPlayerLeft = PrefixPlayer + "Player " + ColorYellow + "%s left the game.\n" + ColorReset
)

// SEC - Server Language
//
//go:embed logo.txt
var Logo string

const Name = ColorCyan + "≡GO" + ColorPurple + "Tower " + ColorReset

const (
	PrefixLog   = ColorCyan + "Log " + ColorPurple + "| " + ColorReset
	PrefixError = ColorRed + "Err " + ColorPurple + "| " + ColorReset
)

const (
	LangServerWelcome = PrefixLog + "Initializing " + Name + "v" + ServerVersion + ".\n"
	LangServerGoodbye = PrefixLog + "Exiting now.\n"
)

const (
	LangTcpListening       = PrefixLog + "TCP Component initialized on TCP:" + ColorYellow + "%d.\n"
	LangTcpClosed          = PrefixLog + "TCP Component exited.\n"
	LangTcpOpenErr         = PrefixError + "TCP Component failed to start. Please ensure the port specified is not in use.\n"
	LangTcpAcceptErr       = PrefixError + "TCP Component failed to accept a connection.\n"
	LangTcpClientConnected = PrefixLog + "TCP Component accepted a connection from " + ColorYellow + "%s" + ColorReset + ".\n"
)

const (
	LangUdpListening = PrefixLog + "UDP Component initialized on UDP:" + ColorYellow + "%d.\n"
	LangUdpClosed    = PrefixLog + "UDP Component exited.\n"
	LangUdpOpenErr   = PrefixError + "UDP Component failed to start. Please ensure the port is not in use.\n"
)

const (
	LangConfigNotfoundErr = PrefixError + "Could not find the config file.\n"
	LangConfigFormatErr   = PrefixError + "Config file is not formatted correctly.\n"
)
