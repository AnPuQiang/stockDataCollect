package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"stockDataCollect/config"
	"stockDataCollect/service"
	"time"

	"github.com/go-zoo/bone"
)

//数据API来源：阿里云 易源数据-股票历史行情查询_免费版

func main() {
	configFile := flag.String("config", "./config.toml", "")
	flag.Parse()
	config.Configuration = config.NewConfig(configFile)
	fmt.Printf("Config is %#v", config.Configuration)
	if len(config.Configuration.StockCode) == 0 {
		panic("if len(Config.StockCode) == 0")
	}

	go service.RoutineQuery()

	mux := bone.New()
	address := fmt.Sprintf(":%d", config.Configuration.Port)
	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Server start serving on", config.Configuration.Port)
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
