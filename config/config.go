package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host           string `json:"service_host"`
	Port           string `json:"service_port"`
	Repository     string `json:"repository"`
	DBPath         string `json:"db_path"`
	MigrationsPath string `json:"migrations_path"`
}

// Holds app configs
var Conf Config

// LoadConfig parses config/config.json file into Config struct.
func LoadConfig() (err error) {
	path := "config/config.json"
	// Get the config file
	var configFile []byte
	configFile, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(configFile, &Conf)
	return
}
