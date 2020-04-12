package service

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

//begin=2015-09-01&code=600004&end=2015-09-02
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

	fmt.Println(string(body))
}
