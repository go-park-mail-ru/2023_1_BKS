package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/config"
	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/internal/app"
)

var (
	configPath = "../../config/config.toml"
)

func main() {
	config := config.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal("Error to decode file: ", err)
	}

	if err := app.Start(config); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on port ", config.BindAddr)
}
