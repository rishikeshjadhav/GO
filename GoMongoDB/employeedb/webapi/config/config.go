package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represent confif file structure
type Config struct {
	Server   string
	Database string
}

// Function to read config file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
