package main

import (
	"log"

	configs "github.com/Lucasmartinsn/grocery-api/Configs"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}
