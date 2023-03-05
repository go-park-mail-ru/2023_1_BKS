package store

type Config struct {
	Host     string `toml:"localhost"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Dbname   string `toml:"dbname"`
}

func NewConfig() *Config {
	return &Config{}
}
