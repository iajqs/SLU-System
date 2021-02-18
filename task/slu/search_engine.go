/**
 * Created by cks
 * Date: 2020-11-25
 * Time: 17:50
 */
package slu

import (

	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"

	"strings"
	"errors"
	// "fmt"
)

const (
	urlBaidu = "http://www.baidu.com"
	userAgent = "Mozilla/5.0 BaiduSpider/2.0"
)

type BaiduSearchProcesser struct {
}

func NewBaiduSearchProcesser() *BaiduSearchProcesser {
	return &BaiduSearchProcesser{}
}

func (this *BaiduSearchProcesser) Finish() {
	return
}


func (this *BaiduSearchProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	
	answer := query.Find(".op_exactqa_s_answer").Text()
	answer = strings.Trim(answer, " \t\n\r")
	if answer == "" {
		answer = query.Find(".op_exactqa_detail_s_answer").Text()
		answer = strings.Trim(answer, " \t\n\r")
	}
	// the entity we want to save by Pipeline
	p.AddField("answer", answer)
}

func SearchEngineAnswer(question string) (string, error) {
	answer := baiduSearch(question)
	
	if answer != "" {
		return answer, nil
	}
	return "", errors.New("empty answer")
}

func baiduSearch(question string) string {
	
	// GetWithParams Params:
	// 1. url
	// 2. Response type is "html" or "json" or "jsonp" or "text"
	// 3. The urltag is name for marking url and distinguish different urls in PageProcesser and Pipeline.
	// 4. The method is POST or GET
	// 5. The postdata is body string sent to server
	// 6. The header is header for http request.
	// 7. Cookies
	spBaidu := spider.NewSpider(NewBaiduSearchProcesser(), "BaiduSearch")
	urlFinal := urlBaidu + "/s?" + "wd=" + question
	// urlFinal = "http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn"
	req := request.NewRequest(urlFinal, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := spBaidu.GetByRequest(req)
	if pageItems == nil {
		return ""
	}
	return pageItems.GetAll()["answer"]
}