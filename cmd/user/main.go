package main

import (
	config "config/user"
	"user"
)

func main() {
	cfg := config.CreateConfig()

	user.Run(cfg)
}
