package Configs

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var cfg *configPostgres

type configPostgres struct {
	DB DBConfigPostgres
}

type DBConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	Driver   string
}

func init() {
	ginMode := os.Getenv("GIN_MODE")
	switch ginMode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func Load() error {
	// Loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading package godoenv: %v", err)
	}

	cfg = new(configPostgres)

	// fetches values ​​from .env file
	cfg.DB = DBConfigPostgres{
		Host:     os.Getenv("DATABASES_HOST"),
		Port:     os.Getenv("DATABASES_PORT"),
		User:     os.Getenv("DATABASES_USER"),
		Pass:     os.Getenv("DATABASES_PASS"),
		Database: os.Getenv("DATABASES_DB"),
		Driver:   os.Getenv("DATABASES_DRIVER"),
	}

	return nil
}

// Returns the struct with values ​​for Database Connection
func GetDB() DBConfigPostgres {
	return cfg.DB
}
