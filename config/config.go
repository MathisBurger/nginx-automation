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

// parses config
func ParseConfig() (c *Config, err error) {
	f, err := os.Open("/root/installation-service/config/config.json")
	if err != nil {
		return
	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	return
}
