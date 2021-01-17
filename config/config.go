package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AllowedOrigins   string `json:"allowed-origins"`
	WebserverAddress string `json:"webserver-address"`
	ApplicationMode  string `json:"application-mode"`
	ApplicationPort  string `json:"application-port"`
}

func ParseConfig() (c *Config, err error) {
	f, err := os.Open("./config/config.json")
	if err != nil {
		return
	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	return
}
