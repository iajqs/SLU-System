/**
 * Created by cks
 * Date: 2020-11-17
 * Time: 11:22
 */
package task

import (
	"SLU-System/task/slu"
	"SLU-System/proto"

	"github.com/sirupsen/logrus"
	
	// "os"
	"time"
	// "fmt"
	"errors"
	"strings"
)

func ReluAnswer(question string) (string, error) {
	return slu.ReluAnswer(question)
}

func ChatAnswer(session, question string) (string, error) {
	return slu.ChatAnswer(session, question)
}

func (task *Task) SearchEngineAnswer(question string) (string, error) {
	return slu.SearchEngineAnswer(question)
}

func SearchEngineAnswer(question string) (string, error) {
	return slu.SearchEngineAnswer(question)
}

func BaikeAnswer(question string) (string, error) {
	return slu.BaikeAnswer(question)
}

func (task *Task) WeatherAnswer(question string) (proto.Weather, error) {
	return slu.WeatherAnswer(question)
}

func WeatherAnswer(question string) (proto.Weather, error) {
	return slu.WeatherAnswer(question)
}

func MusicAnswer(question string) (string, error) {
	return slu.MusicAnswer(question)
}

func ModelAnswer(question string) (string, string, error) {
	return slu.ModelAnswer(question)
}

func (task *Task) SLUAnswer(question string) (string, string, error) {
	question = strings.Trim(question, " \t\n\r")
	cReluAnswer := make(chan string, 1)
	cChatAnswer := make(chan string, 1)
	cSearchEngineAnswer := make(chan string, 1)
	cBaikeAnswer := make(chan string, 1)
	cWeatherAnswer := make(chan proto.Weather, 1)
	cMusicAnswer := make(chan string, 1)
	cModelIntentAnswer := make(chan string, 1)
	cModelKeywordsAnswer := make(chan string, 1)
	cEnd := make(chan int, 1)
	go func(question string, cReluAnswer chan string) {
		answer, _ := ReluAnswer(question)
		// if err != nil {
		// 	logrus.Info("规则出错: ", err.Error())
		// }
		cReluAnswer <- answer
	} (question, cReluAnswer)

	go func(question string, cChatAnswer chan string) {
		answer, _ := ChatAnswer("10001", question)
		// if err != nil {
		// 	logrus.Info("闲聊出错: ", , err.Error())
		// }
		cChatAnswer <- answer
	} (question, cChatAnswer)

	go func(question string, cSearchEngineAnswer chan string) {
		answer, _ := SearchEngineAnswer(question)
		// if err != nil {
			// logrus.Info("百度搜索出错: ", , err.Error())
		// }
		cSearchEngineAnswer <- answer
	} (question, cSearchEngineAnswer)

	go func(question string, cBaikeAnswer chan string) {
		answer, _ := BaikeAnswer(question)
		// if err != nil {
		// 	logrus.Info("百科出错: ", , err.Error())
		// }
		cBaikeAnswer <- answer
	} (question, cBaikeAnswer)

	go func(question string, cWeatherAnswer chan proto.Weather) {
		answer, _ := WeatherAnswer(question)
		// if err != nil {
		// 	logrus.Info("天气查询出错: ", , err.Error())
		// }
		cWeatherAnswer <- answer
	} (question, cWeatherAnswer)

	go func(question string, cMusicAnswer chan string) {
		answer, _ := MusicAnswer(question)
		// if err != nil {
		// 	logrus.Info("音乐搜索出错: ", , err.Error())
		// }
		cMusicAnswer <- answer
	} (question, cMusicAnswer)

	go func(question string, cModelIntentAnswer chan string, cModelKeywordsAnswer chan string) {
		intent, keywords, err := ModelAnswer(question)
		if err != nil {
			logrus.Infof("模型出错: ", err.Error())
		}
		cModelIntentAnswer <- intent
		cModelKeywordsAnswer <- keywords
	} (question, cModelIntentAnswer, cModelKeywordsAnswer)

	now := int(time.Now().Unix())
	go func(now int, cEnd chan int) {
		for {
			if int(time.Now().Unix()) - now >= 2 {
				cEnd <- 1
				return
			}
		}
	}(now, cEnd)

	// relu := ""
	chat := ""
	baidu := ""
	baike := ""
	weather := proto.Weather{}
	music := ""
	model := ""
	modelKeywords := ""
	for {
		select {
		// case relu = <- cReluAnswer:
			// fmt.Println("relu answer :", relu)
		case chat = <- cChatAnswer:
			// fmt.Println("chat answer :", chat)
		case baidu = <- cSearchEngineAnswer:
			// fmt.Println("baidu answer :", baidu)
		case baike = <- cBaikeAnswer:
			// fmt.Println("baike answer :", baike)
		case weather = <- cWeatherAnswer:
			// fmt.Println("weather answer :", weather.ToString())
		case music = <- cMusicAnswer:
			// fmt.Println("music answer :", music)
		case model = <- cModelIntentAnswer:
			// fmt.Println("model answer :", model)
		case modelKeywords = <- cModelKeywordsAnswer:

		case <- cEnd:
			baike = "" // 移除百科查询结果
			if model != "" {
				if model == "weather" &&  weather.Code != 404 && weather.TemperatureNow != ""{
					return weather.ToString(), "weather", nil
				} else if model == "music"{
					if modelKeywords != ""{
						music, _ = MusicAnswer(modelKeywords)
						return music, "music", nil
					}
					return music, "music", nil
				} else if model == "poetry" && baidu != "" {
					return baidu, "baidu", nil
				} else if chat != "" {
					return chat, "chat", nil
				}
			} else {
				if weather.Code != 404 && weather.TemperatureNow != ""{
					return weather.ToString(), "weather", nil
				} else if baidu != "" {
					return baidu, "baidu", nil
				} else if baike != "" {
					return baike, "baike", nil
				} else if chat != "" {
					return chat, "chat", nil
				}
			}
			return "", "", errors.New("answer empty")
		}
	}
	return "", "", errors.New("answer empty")
}