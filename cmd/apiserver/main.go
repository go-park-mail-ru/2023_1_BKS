package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal("Error to decode file: ", err)
	}
	// Запускаем сервер
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
