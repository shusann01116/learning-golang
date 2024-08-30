package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      uint16 `envconfig:"PORT" default:"3000"`
	Host      string `envconfig:"HOST" required:"true"`
	AdminPort uint16 `envconfig:"ADMIN_PORT" default:"3001"`
}

func main() {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(err)
		return
	}
	log.Println(c)
}
