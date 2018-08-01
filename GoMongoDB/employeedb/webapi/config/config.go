package config

import (
	"fmt"
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
	fmt.Printf("\nReading from file %s\n", "config.toml")
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
