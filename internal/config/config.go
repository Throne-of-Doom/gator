package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	path := filepath.Join(home, ".gatorconfig.json")
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	var cfg Config
	err = json.NewDecoder(file).Decode(&cfg)
	return cfg, err
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, ".gatorconfig.json")

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0o600)
}
