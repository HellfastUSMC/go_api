package store

type Config struct {
	DBURL string `toml:"database_url"`
}

func NewConf() *Config {
	return &Config{}
}
