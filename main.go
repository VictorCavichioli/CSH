package main

import (
	"fmt"
	"log"
	"csh-api/cmd/config"
	"csh-api/cmd/rest"
)

const configPath string = "./config/csh.yaml"

func main() {
	loadConfig()
	rest.StartServer()
}

func loadConfig() {
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	} else {
		fmt.Printf("Loaded Server Config: %+v\n", config)
	}
}