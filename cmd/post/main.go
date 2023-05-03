package main

import (
	"post"
	"post/config"
)

func main() {
	cfg := config.CreateConfig()

	post.Run(cfg)
}
