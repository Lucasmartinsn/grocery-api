package confDB

import (
	"log"
	"os"

	load "github.com/Lucasmartinsn/grocery-api/Configs/confEnv"
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
	// Loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading package godoenv: %v", err)
	}
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

func loadDB() error {
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

func Load() error {
	if errdb := loadDB(); errdb != nil {
		return errdb
	} else if errev := load.LoadEnv(); errev != nil {
		return errev
	} else {
		return nil
	}
}

// Returns the struct with values ​​for Database Connection
func GetDB() DBConfigPostgres {
	return cfg.DB
}
