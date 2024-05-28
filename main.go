package main

import (
	"log"

	configs "github.com/Lucasmartinsn/grocery-api/Configs"
	"github.com/Lucasmartinsn/grocery-api/Server"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	server := Server.NewServer()
	server.Run()
}
