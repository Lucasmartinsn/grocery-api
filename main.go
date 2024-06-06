package main

import (
	"log"

	configs "github.com/Lucasmartinsn/grocery-api/Configs/confDB"
	"github.com/Lucasmartinsn/grocery-api/Server"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("error loading Server: %v", err)
	}

	server := Server.NewServer()
	server.Run()
}
