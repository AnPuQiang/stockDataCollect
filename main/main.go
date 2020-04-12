package main

import "stockDataCollect/service"

//数据API来源：阿里云 易源数据-股票历史行情查询_免费版
func main() {
	service.GetStockData("2020-04-01", "2020-04-02", "600004")
}
