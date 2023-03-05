package apiserver

import "github.com/go-park-mail-ru/2023_1_BKS/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` // Адрес запуска сервера
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
