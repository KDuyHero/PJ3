package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Source        string
	MigrateFolder string
}

type Config struct {
	PostGres DBConfig
}

func LoadDBConfig() DBConfig {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	source := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	fmt.Println(source)
	return DBConfig{
		Source:        source,
		MigrateFolder: os.Getenv("DB_MIGRATION_FOLDER"),
	}
}

func Load() *Config {
	return &Config{
		PostGres: LoadDBConfig(),
	}
}
