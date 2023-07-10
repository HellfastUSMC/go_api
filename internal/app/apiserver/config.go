package apiserver

import "github.com/HellfastUSMC/goapi/internal/app/store"

// Config type
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLVL   string `toml:"log_lvl"`
	Store    *store.Config
}

// NewConfig API Server config constructor
func NewConfig() *Config {
	return &Config{
		BindAddr: ":6060",
		LogLVL:   "debug",
		Store:    store.NewConf(),
	}
}
