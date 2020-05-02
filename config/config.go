package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Port      int      `toml:"port"`
	AppCode   string   `toml:"app_code"`
	StockCode []string `toml:"stock_code"`
}

var Configuration *Config

func NewConfig(configPath *string) *Config {
	var stockConfig Config
	fmt.Println(*configPath)
	if _, err := toml.DecodeFile(*configPath, &stockConfig); err != nil {
		panic(err)
	}

	//if stockConfig.AppCode == "" {
	//	stockConfig.AppCode = "82399585b13d43b49a42368654e5bd44"
	//}
	//if len(stockConfig.StockCode) == 0{
	//	stockConfig.StockCode = append(stockConfig.StockCode, "300750", "000100")
	//}
	return &stockConfig
}
