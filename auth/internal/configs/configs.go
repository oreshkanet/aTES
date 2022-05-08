package configs

import (
	"os"
)

type Configs struct {
	Port       string
	SigningKey string
	HashSalt   string
}

func Load() *Configs {
	return &Configs{
		Port:       os.Getenv("PORT"),
		SigningKey: os.Getenv("SIGNING_KEY"),
		HashSalt:   os.Getenv("HASH_SALT"),
	}
}
