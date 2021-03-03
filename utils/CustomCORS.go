package utils

import "github.com/MathisBurger/nginx-automation/config"

// checks custom cors
func CheckCORS(ip string) bool {
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}
	return cfg.AllowedOrigins == ip
}
