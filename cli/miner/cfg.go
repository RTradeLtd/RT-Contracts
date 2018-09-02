package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config holds the configuration for our miner
type Config struct {
	Endpoint        string `json:"endpoint"`
	KeyFilePath     string `json:"key_file_path"`
	KeyFilePass     string `json:"key_file_pass"`
	ContractAddress string `json:"contract_address"`
}

// LoadConfig is used to load our configuration
func LoadConfig(configPath string) (*Config, error) {
	var cfg Config
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
