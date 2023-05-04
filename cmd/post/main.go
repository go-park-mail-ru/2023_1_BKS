package main

import (
	config "github.com/go-park-mail-ru/2023_1_BKS/config/post"
	"github.com/go-park-mail-ru/2023_1_BKS/internal/post"
)

func main() {
	cfg := config.CreateConfig()

	post.Run(cfg)
}
