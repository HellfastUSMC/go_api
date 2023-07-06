package apiserver

//Config type
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLVL   string `toml:"log_lvl"`
}

//New Config constructor
func NewConfig() *Config {
	return &Config{}
}
