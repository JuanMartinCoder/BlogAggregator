package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/JuanMartinCoder/BlogAggregator/internal/database"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBurl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
type State struct {
	Cfg *Config
	DB  *database.Queries
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(config Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}
