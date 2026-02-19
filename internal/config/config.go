package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)
func Read() (Config,error){
	homeDir,err := os.UserHomeDir()

	if err != nil{
		return Config{},err
	}
	
	path := filepath.Join(homeDir,configFileName) 
	content, err := os.ReadFile(path)

	if err != nil{
		return Config{},err
	}

	var conf Config
	err = json.Unmarshal(content,&conf)

	if err != nil{
		return Config{},err
	}
	
	return conf,nil	
}
