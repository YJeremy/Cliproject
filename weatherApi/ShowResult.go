package weatherApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//天气数据处理，及其界面显示
func Print(day string, r Response) {
	fmt.Println("城市：", r.CityInfo.City)
	if day == "今天" {
		fmt.Println("日期：", r.Date)
		fmt.Println("湿度：", r.Data.ShiDu)
		fmt.Println("空气质量：", r.Data.Quality)
		fmt.Println("温馨提示：", r.Data.Ganmao)
		fmt.Println("温度:", r.Data.Wendu, "。C")
		fmt.Println("PM2.5:", r.Data.Pm25)
		fmt.Println("PM10:", r.Data.Pm10)
	} else if day == "昨天" {
		fmt.Println("日期：", r.Data.Yestoday.Date)
		fmt.Println("温度：", r.Data.Yestoday.Low, r.Data.Yestoday.High)
		fmt.Println("风量：", r.Data.Yestoday.Fx, r.Data.Yestoday.F1)
		fmt.Println("天气：", r.Data.Yestoday.Type)
		fmt.Println("温馨提示:", r.Data.Yestoday.Notice)
	} else if day == "预测" {
		fmt.Println("=======================================")
		for _, item := range r.Data.Forecast {
			fmt.Println("日期：", item.Date)
			fmt.Println("温度：", item.Low, item.High)
			fmt.Println("风量：", item.Fx, item.F1)
			fmt.Println("天气：", item.Type)
			fmt.Println("温馨提示：", item.Notice)
			fmt.Println("=========================")
		}
	} else {
		fmt.Println("...")
	}
}

//查询天气功能数据解析

func Weather(city, day string) error {

	const apiURL = "http://t.weather.sojson.com/api/weather/city/"
	var body, err = Requests(apiURL + city) //解析目标网址的 API
	if err != nil {
		fmt.Printf("err was %v", err)
		return nil
	}

	//创建变量，存放获取API被解析后的数据
	var r Response
	err = json.Unmarshal([]byte(body), &r) //解析数据 从URL获取的body转化为[]byte, 再json格式解析，解析后放入类型r
	if err != nil {
		fmt.Printf("\nError message: %v", err)
		return nil
	}
	if r.Status != 200 {
		fmt.Printf("获取天气API出现错误啦，%s", r.Message)
		return nil
	}

	//在解析后的数据里，选择需要的数据显示出来
	Print(day, r)
	return nil

}

//网址解析函数
func Requests(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}
