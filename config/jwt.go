package config

import (
	"log"
	"os"
)

var JWTSecret []byte

func LoadJWTSecret() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is missing")
	}

	JWTSecret = []byte(secret)
}
