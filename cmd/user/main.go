package main

import (
	"github.com/go-park-mail-ru/2023_1_BKS/internal/user"

	config "github.com/go-park-mail-ru/2023_1_BKS/config/user"
)

func main() {
	cfg := config.CreateConfig()

	user.Run(cfg)
}
