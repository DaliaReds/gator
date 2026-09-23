package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB   string `json:"db_url"`
	User string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	file, err := os.Open(configFileName)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(user string) error {
	c.User = user

	file, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(c)
}
