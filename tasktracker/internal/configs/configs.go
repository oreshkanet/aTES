package configs

import (
	"os"
)

type Configs struct {
	Port       string
	SigningKey string

	MsSqlhost string
	MsSqlDb   string
	MsSqlUser string
	MsSqlPwd  string

	KafkaHost string
	KafkaPort string
}

func Load() *Configs {
	return &Configs{
		Port:       os.Getenv("PORT"),
		SigningKey: os.Getenv("SIGNING_KEY"),

		MsSqlhost: os.Getenv("MSSQL_HOST"),
		MsSqlDb:   os.Getenv("MSSQL_DB"),
		MsSqlUser: os.Getenv("MSSQL_USER"),
		MsSqlPwd:  os.Getenv("MSSQL_PWD"),

		KafkaHost: os.Getenv("KAFKA_HOST"),
		KafkaPort: os.Getenv("KAFKA_PORT"),
	}
}
