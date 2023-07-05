package main

import (
	"1projects/go_api/internal/app/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/HellfastUSMC/goapi"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := go_api.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
