/**
 * Created by cks
 * Date: 2020-11-26
 * Time: 15:49
 */
package slu

import (
	"github.com/bitly/go-simplejson"

	"net/url"
	"net/http"
	"io/ioutil"
	"errors"
)

var (
	urlModel = "http://127.0.0.1:5000/nlp"
	client = &http.Client{}
)

func ModelAnswer(question string) (string, string, error) {
	paras := url.Values{"question": {question}}
	urlReal := urlModel + "?" + paras.Encode()
	resp, err := http.Get(urlReal)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	
	js, err := simplejson.NewJson(body)
	if err != nil {
		return "", "", errors.New("json 解析失败")
	}
	intent, err := js.Get("intent").String()
	if err != nil {
		return "", "", errors.New("没有意图结果")
	}
	keywords, err := js.Get("keywords").String()
	return intent, keywords, nil
}