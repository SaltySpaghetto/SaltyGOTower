package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

// Config is simply a struct to represent a config file's options.
type Config struct {
	TCPPort uint
	UDPPort uint

	MaxPlayers uint
}

// LoadConfig loads a config file at the specified path and returns a struct containing its data.
func LoadConfig(path string) Config {
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf(LangConfigNotfoundErr)
		os.Exit(-1)
	}

	var conf Config

	// SEC - Networking Values
	conf.TCPPort, err = cfg.Section("networking").Key("tcp_port").Uint()
	conf.UDPPort, err = cfg.Section("networking").Key("udp_port").Uint()

	// SEC - Administration Values
	conf.MaxPlayers, err = cfg.Section("administration").Key("max_players").Uint()

	if err != nil {
		fmt.Printf(LangConfigFormatErr)
		os.Exit(1)
	}

	return conf
}
