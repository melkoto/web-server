package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/melkoto/web-server/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {

	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	// В этот момент происходит инициализация переменной configPath
	flag.Parse()
	fmt.Println("It works")

	// server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Panicln("can't find config file. using default values:", err)
	}

	server := api.New(config)

	// start api server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	//log.Fatal(server.Start())
}
