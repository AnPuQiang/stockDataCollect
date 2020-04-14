package main

import (
	"flag"
	"fmt"
	"stockDataCollect/config"
	"stockDataCollect/service"
	"time"

	log "github.com/Sirupsen/logrus"
)

//数据API来源：阿里云 易源数据-股票历史行情查询_免费版

var Config *config.Config

func main() {
	configFile := flag.String("config", "./config.toml", "")
	flag.Parse()
	Config = config.NewConfig(configFile)
	fmt.Printf("Config is %#v", Config)
	const (
		updateIntervalHour = 1
		updateBeginMinute  = 5
	)
	loc, _ := time.LoadLocation("Asia/Chongqing")

	for {
		now := time.Now()
		next := now.Add(time.Hour * updateIntervalHour)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), updateBeginMinute, 0, 0, loc)
		log.Infof("the interval is ", next.Sub(now).String())
		timer := time.NewTimer(next.Sub(now))

		format := time.Now().Format("2006-01-01")
		fmt.Println("format:", format)
		<-timer.C
		if next.Hour() > 15 || next.Hour() < 9 {
			continue
		}

		if len(Config.StockCode) == 0 {
			panic("if len(Config.StockCode) == 0")
		}
		for i := 0; i < len(Config.StockCode); i++ {
			service.GetStockData(format, format, Config.StockCode[i])

		}
		time.Sleep(time.Minute)
	}

}
