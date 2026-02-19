package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct{
	Username string `json:"current_user_name"`
	Url string `json:"db_url"`
}

func (c *Config) SetUser(username string) error{
	c.Username = username
	
	data,err := json.Marshal(c)
	if err != nil {
		return err
	}
	
	home,err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, configFileName)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}	
