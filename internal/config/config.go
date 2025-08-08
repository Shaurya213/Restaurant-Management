package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	JWTSecret string
	JWTExpiry time.Duration
}

func Load() *Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("JWT_SECRET not set, using dev default (CHANGE IN PROD)")
		secret = "dev_secret_change_me"
	}

	expStr := os.Getenv("JWT_EXP_MINUTES")
	if expStr == "" {
		expStr = "120"
	}
	mins, err := strconv.Atoi(expStr)
	if err != nil || mins <= 0 {
		mins = 120
	}

	return &Config{
		JWTSecret: secret,
		JWTExpiry: time.Duration(mins) * time.Minute,
	}
}
