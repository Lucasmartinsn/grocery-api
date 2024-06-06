package confEnv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var cfEnv *configEnv

type configEnv struct {
	Env EncConf
}

type EncConf struct {
	Env string
}

func LoadEnv() error {
	// Loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading package godoenv: %v", err)
	}

	cfEnv = new(configEnv)

	cfEnv.Env = EncConf{
		Env: os.Getenv("AES_SECRET_KEY"),
	}
	return nil
}

// Returns the struct with values ​​for key Connection
func getENV() EncConf {
	return cfEnv.Env
}

func Variable() string {
	env := getENV()
	return env.Env
}
