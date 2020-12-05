/**
 * Created by cks
 * Date: 2020-11-25
 * Time: 15:52
 */
package testtask

import (
	"SLU-System/task"
	"testing"
	// "strings"
)

var tk = task.New()

var questions = []string {
	"你好",
	"你今年几岁啦",
	"姚明的身高",
	"床前明月光",
	"定一个两点的闹钟",
	"上海今天的天气咋样",
	"早上好",
	"你中午吃什么",
	"珠穆朗玛峰有多高",
}


// func TestRelu(t *testing.T) {
// 	for _, question := range questions {
// 		answer, err := tk.ReluAnswer(question)
// 		if err != nil {
// 			t.Error(err.Error())
// 		}else {
// 			t.Log("question: ", question, "answer: ", answer)
// 		}
// 	}
// }

// func TestChat(t *testing.T) {
// 	for _, question := range questions {
// 		answer, err := tk.ChatAnswer("10001", question)
// 		if err != nil {
// 			t.Error(err.Error())
// 		}else {
// 			t.Log("question: ", question, "answer: ", answer)
// 		}
// 	}
// }

func TestSearchEngine(t *testing.T) {
	for _, question := range questions {
		answer, err := tk.SearchEngineAnswer(question)
		if err != nil { 
			t.Error(err.Error())
		}else {
			t.Log("question: ", question, "answer: ", answer)
		}
	}
}

// func TestBaike(t *testing.T) {
// 	for _, question := range questions {
// 		answer, err := tk.BaikeAnswer(question)
// 		if err != nil { 
// 			t.Error(err.Error())
// 		} else {
// 			t.Log("question: ", question, "answer: ", answer)
// 		}
// 	}
// }

func TestWeather(t *testing.T) {
	for _, question := range questions {
		answer, err := tk.WeatherAnswer(question)
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Log("question: ", question, "answer: ", answer.ToString())
		}
	}
}

// func TestMusic(t *testing.T) {
// 	for _, question := range questions {
// 		answer, err := tk.MusicAnswer(question)
// 		if err != nil {
// 			t.Error(err.Error())
// 		} else {
// 			t.Log("question: ", question, "answer: ", answer)
// 		}
// 	}
// }

// func TestModel(t *testing.T) {
// 	for _, question := range questions {
// 		intent, keywords, err := tk.ModelAnswer(question)
// 		if err != nil {
// 			t.Error(err.Error())
// 		} else {
// 			t.Log("question: ", question)
// 			t.Log("question: ", intent)
// 			t.Log("keywords: ", strings.Join(keywords, ","))
// 		}
// 	}
// }

func TestSLU(t *testing.T) {
	for _, question := range questions {
		answer, source, err := tk.SLUAnswer(question)
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Log("question: ", question, "answer: ", answer, "source:", source)
		}
	}
}