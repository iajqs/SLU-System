/**
 * Created by cks
 * Date: 2020-11-25
 * Time: 16:00
 */
package slu

import (

)


import (
	"SLU-System/proto"

	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"

	"strings"
	"errors"
	"strconv"
	// "fmt"
)

const (
	urlWeather = "http://www.baidu.com"
	// userAgent = "Mozilla/5.0 BaiduSpider/2.0"
)

type WeatherProcesser struct {
}

func NewWeatherProcesser() *WeatherProcesser {
	return &WeatherProcesser{}
}

func (this *WeatherProcesser) Finish() {
	return
}

var (
	spWeather = spider.NewSpider(NewWeatherProcesser(), "Weather")
)

func (this *WeatherProcesser) Process(p *page.Page) {
	
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	code := "0"
	query := p.GetHtmlParser()
	
	title := query.Find(".t.c-gap-bottom-small a").Text()
	// fmt.Printf(query.Find(".t.c-gap-bottom-small a").Text())
	title = strings.Trim(title, " \t\n\r")

	weatherToday := query.Find(".op_weather4_twoicon_today.OP_LOG_LINK")

	if strings.Trim(weatherToday.Text(), " \t\n\r") == "" {
		code = "404"
	}
	temperatureNow 	 := strings.Trim(weatherToday.Find(".op_weather4_twoicon_shishi_title").Text(), " \t\n\r")
	temperatureRange := strings.Trim(weatherToday.Find(".op_weather4_twoicon_temp").Text(), " \t\n\r")
	weatherNow 		 := strings.Trim(weatherToday.Find(".op_weather4_twoicon_shishi_sub").Text(), " \t\n\r")
	weatherPredict 	 := strings.Trim(weatherToday.Find(".op_weather4_twoicon_weath").Text(), " \t\n\r")
	wind 			 := strings.Trim(weatherToday.Find(".op_weather4_twoicon_wind").Text(), " \t\n\r")
	airQuality 		 := strings.Trim(weatherToday.Find(".op_weather4_twoicon_realtime_quality_today").Text(), " \t\n\r")
	// weatherToday = strings.Trim(weatherToday)
	// the entity we want to save by Pipeline
	p.AddField("code", code)
	p.AddField("title", title)
	p.AddField("temperatureNow", temperatureNow)
	p.AddField("temperatureRange", temperatureRange)
	p.AddField("weatherNow", weatherNow)
	p.AddField("weatherPredict", weatherPredict)
	p.AddField("wind", wind)
	p.AddField("airQuality", airQuality)
}

func WeatherAnswer(question string) (proto.Weather, error) {
	answer, err := weatherSearch(question)
	if err != nil {
		return proto.Weather{}, err
	}
	if answer.Code == 0 {
		return answer, nil
	}
	return proto.Weather{Code: 404}, errors.New("empty answer")
}


func weatherSearch(question string) (proto.Weather, error) {
	urlFinal := urlBaidu + "/s?" + "wd=" + question

	req := request.NewRequest(urlFinal, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := spWeather.GetByRequest(req)

	result := pageItems.GetAll()
	
	code, err := strconv.Atoi(result["code"])
	if err != nil {
		return proto.Weather{}, err
	}
	return proto.Weather {
		Code:				code,
		TemperatureNow: 	result["temperatureNow"],
		TemperatureRange:   result["temperatureRange"],
		WeatherNow: 		result["weatherNow"],
		WeatherPredict: 	result["weatherPredict"],
		Wind: 			 	result["wind"],
		AirQuality: 		result["airQuality"],
	}, nil
}