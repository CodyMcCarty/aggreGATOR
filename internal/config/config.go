package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	path += "/.gatorconfig.json"
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg Config) SetUser(UserName string) error {
	cfg.CurrentUserName = UserName

	path, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path += "/.gatorconfig.json"

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0666)
	if err != nil {
		return err
	}

	return nil
}
