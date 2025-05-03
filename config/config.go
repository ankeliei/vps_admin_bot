package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Server struct {
		Addr     string `json:"addr"`
		Host     string `json:"host"`
		CertFile string `json:"certFile"`
		KeyFile  string `json:"keyFile"`
	} `json:"server"`
}

func LoadJSONConfig(path string) (*config, error) {
	fmt.Print(path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析失败: %v", err)
	}
	fmt.Print(config)
	return &config, nil
}
