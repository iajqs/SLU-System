/**
 * Created by cks
 * Date: 2020-11-26
 * Time: 19:49
 */
package slu

import (
	
	"github.com/bitly/go-simplejson"
	"net/http"
	"net/url"
	"fmt"
	// "bytes"
	"io/ioutil"
	"encoding/json"
	"errors"
)

var (
	urlMusic = "http://localhost:3000"
	// client = &http.Client{}
)

func MusicAnswer(question string) (string, error) {
	paras := url.Values{"keywords": {question}}
	urlReal := urlMusic + "/search?" + paras.Encode()
	resp, err := http.Get(urlReal)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	
	js, err := simplejson.NewJson(body)
	if err != nil {
		return "", errors.New("json 解析失败")
	}
	var nodes = make(map[string]interface{})
	nodes, err = js.Get("result").Get("songs").GetIndex(0).Map()
	if err != nil {
		return "", errors.New("音乐搜索没有结果")
	}
	urlResult := fmt.Sprintf("http://music.163.com/song/media/outer/url?id=%s.mp3", nodes["id"].(json.Number))
	return urlResult, nil
}