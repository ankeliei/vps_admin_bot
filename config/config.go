package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Server struct {
		Addr     string `json:"addr"`
		CertFile string `json:"certFile"`
		KeyFile  string `json:"keyFile"`
	} `json:"server"`
	Telebot struct {
		BotApi  string `json:"botApi"`
		Debug   bool   `json:"debug"`
		Timeout int    `json:"timeout"`
	} `json:"telebot"`
}

var (
	ServerCFG *config
)

func loadJSONConfig(path string) (*config, error) {
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

func init() {
	var err error
	ServerCFG, err = loadJSONConfig("/home/ank/workspace/vps_admin_bot/config/config.json")
	if err != nil {
		panic(err)
	}
}
