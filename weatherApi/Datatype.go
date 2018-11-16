package weatherApi

//需要根据API网址返回的JSON的标号定义字段 http://t.weather.sojson.com/api/weather/city/
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	time    string `json:"time"`

	CityInfo city   `json:"cityInfo"`
	Date     string `json:"date"`
	Data     Data   `json:"data"` //自定义的数据类型，见下
}

type city struct {
	City       string `json:"city"`
	Parent     string `json:"parent"`
	UpdateTime string `json:"updateTime"`
	CityId     string `json:"cityId"`
}

type Data struct {
	ShiDu    string  `json:"shidu"`
	Pm25     float32 `json:"pm25"`
	Pm10     float32 `json:"pm10"`
	Quality  string  `json:"quality"`
	Ganmao   string  `json:"ganmao"`
	Wendu    string  `json:"wendu"`
	Yestoday Day     `json:"yestoday"`
	Forecast []Day   `json:"forecast"` //自定义的数据类型，见下;未来5天,所以有多个同类型组成的切片
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
