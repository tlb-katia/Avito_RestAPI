package openapi

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress    string
	PostgresConn     string
	PostgresJdbcUrl  string
	PostgresUsername string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
}

func MustLoad() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	return &Config{
		ServerAddress:    os.Getenv("SERVER_ADDRESS"),
		PostgresConn:     os.Getenv("POSTGRES_CONN"),
		PostgresJdbcUrl:  os.Getenv("POSTGRES_JDBC_URL"),
		PostgresUsername: os.Getenv("POSTGRES_USERNAME"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE=tender-service\n"),
	}
}
