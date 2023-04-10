package main

import (
	"config"
	"user"
)

func main() {
	cfg := config.CreateConfig()

	user.Run(cfg)
}
