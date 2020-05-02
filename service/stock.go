package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"stockDataCollect/config"
	"time"

	log "github.com/Sirupsen/logrus"
)

func RoutineQuery() {

	for {
		now := time.Now()

		format := time.Now().Format("2006-01-01")
		fmt.Println("format:", format)
		fmt.Println("today is ", now.Weekday())
		//filter time
		if now.Weekday() == time.Sunday || now.Weekday() == time.Saturday {
			time.Sleep(time.Hour)
			continue
		}
		//9:30-10:00 10:00-11:00 11:00-11:30 && 13:00-15:00
		if (now.Hour() == 9 && now.Minute() >= 30) || (now.Hour() == 10) || (now.Hour() == 11 && now.Minute() <= 30) ||
			(now.Hour() == 13) || (now.Hour() == 14) {

			for i := 0; i < len(config.Configuration.StockCode); i++ {
				GetStockData(format, format, config.Configuration.StockCode[i])
			}

		} else {
			time.Sleep(time.Minute)

			continue
		}

		time.Sleep(time.Minute)

	}

}

//begin=2015-09-01&code=600004&end=2015-09-02
/** response format
{
    "showapi_res_error":"",
    "showapi_res_id":"f3b49de402564d79a30c415c31c6b8e0",
    "showapi_res_code":0,
    "showapi_res_body":{
        "ret_code":0,
        "showapi_fee_code":-1,
        "list":[
            {
                "trade_money":"135201992", //交易金额
                "diff_money":"0.20",	//涨跌金额
                "code":"600004",		//股票代码
                "open_price":"12.46",	//开盘价
                "date":"2020-04-02",	//日期
                "market":"sh",			//市场
                "min_price":"12.27",	//最低价
                "trade_num":"108310",	//交易手数
                "turnover":"0.52",		//换手率
                "close_price":"12.65",	//收盘价
                "max_price":"12.67",	//最高价
                "swing":"3.21",			//振幅
                "diff_rate":"1.61"		//涨跌幅(%)
            },
            {
                "trade_money":"105467835",
                "diff_money":"-0.10",
                "code":"600004",
                "open_price":"12.52",
                "date":"2020-04-01",
                "market":"sh",
                "min_price":"12.45",
                "trade_num":"83741",
                "turnover":"0.4",
                "close_price":"12.45",
                "max_price":"12.81",
                "swing":"2.87",
                "diff_rate":"-0.8"
            }
        ]
    }
}
*/
func GetStockData(begin string, end string, code string) {
	host := "http://stock.market.alicloudapi.com"
	path := "/sz-sh-stock-history"
	appCode := "82399585b13d43b49a42368654e5bd44"
	query := "?begin=" + begin + "&code=" + code + "&end=" + end
	url := host + path + query
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "APPCODE "+appCode)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("[GetStockData] client.Do(req) err is", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("[GetStockData] ioutil.ReadAll(resp.Body) err is", err.Error())
	}

	//Decode json and save in ES, ES in localhost
	//if the stock is volatile, send it through wechat
	fmt.Println(string(body))
}
