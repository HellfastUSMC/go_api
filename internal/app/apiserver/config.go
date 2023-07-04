package apiserver

//Config type
type Config struct {
	BindAddr string `toml: "bind_addr"`
}

//New Config constructor
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
