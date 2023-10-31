package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Scaler struct {
	SerialPort string `json:"SerialPort"`
	Type       string `json:"Type"`
	Protocol   string `json:"Protocol"`
}

type GRPC struct {
	Address string `json:"Address"`
}

type Log struct {
	LogLevel string `json:"Level"`
}

type Config struct {
	Scaler `json:"Port"`
	GRPC   `json:"gRPCServer"`
	Log    `json:"Log"`
}

func New() (*Config, error) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer jsonFile.Close()

	content, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	cfg := &Config{}
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	
	return cfg, nil
}
