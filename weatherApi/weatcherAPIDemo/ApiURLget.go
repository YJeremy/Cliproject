package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Status   int    `json:"status"`
	CityName string `json:"city"`
	Data     Date   `json:"data"`
	Date     string `json:"date"`
	Message  string `json:"message"`
	Count    int    `json:"count"`
}

type Date struct {
	ShiDu    string `json:"shidu"`
	Quality  string `json:"quality"`
	Ganmao   string `json:"ganmao"`
	Yestoday Day    `json:"yestoday"`
	Forecast []Day  `json:"forecast"`
}

type Day struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"Sunset"`
	Aqi     float32 `json:"aqi"`
	Fx      string  `json:"fx"`
	F1      string  `json:"f1"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}

func main() {
	const apiURL = "http://www.sojson.com/open/api/weather/json.shtml?city="
	app := cli.NewApp()
	app.Name = "weather-Jeremy"
	app.Usage = "天气小预报程序"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "city,c",
			Value: "广州",
			Usage: "城市中文名",
		},
		cli.StringFlag{
			Name:  "day,d",
			Value: "今天",
			Usage: "可选：今天，昨天，预测",
		},
		cli.StringFlag{
			Name:  "Author, r",
			Value: "Jeremy",
			Usage: "Author name",
		},
	}

	app.Action = func(c *cli.Context) error {
		city := c.String("city")
		day := c.String("day")

		var body, err = Requests(apiURL + city) //网址的 API格式
		if err != nil {
			fmt.Printf("err was %v", err)
			return nil
		}

		var r Response
		err = json.Unmarshal([]byte(body), &r) // 从URL获取的body转化为[]byte, 再json格式解析，解析后放入类型r
		if err != nil {
			fmt.Printf("\nError message: %v", err)
			return nil
		}
		if r.Status != 200 {
			fmt.Printf("获取天气API出现错误啦，%s", r.Message)
			return nil
		}
		Print(day, r)
		return nil
	}
	app.Run(os.Args)
}

func Print(day string, r Response) {
	fmt.Println("城市：", r.CityName)
	if day == "今天" {
		fmt.Println("湿度：", r.Data.ShiDu)
		fmt.Println("空气质量：", r.Data.Quality)
		fmt.Println("温馨提示：", r.Data.Ganmao)
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

func Requests(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

//求两个数最大公约数
func gbs(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gbs(b, a%b)
	}
}

//求多个数的最大公约数
