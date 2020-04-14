package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	AppCode   string   `json:"app_code"`
	StockCode []string `json:"stock_code"`
}

func NewConfig(configPath *string) *Config {
	var stockConfig Config
	fmt.Println(*configPath)
	if _, err := toml.DecodeFile(*configPath, &stockConfig); err != nil {
		panic(err)
	}
	return &stockConfig
}
