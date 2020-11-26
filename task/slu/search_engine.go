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
	urlBaike = "http://baike.baidu.com"
	userAgent = "Mozilla/5.0 BaiduSpider/2.0"
)

type BaiduSearchProcesser struct {
}

type BaikeSearchProcesser struct {
}

func NewBaiduSearchProcesser() *BaiduSearchProcesser {
	return &BaiduSearchProcesser{}
}

func NewBaikeSearchProcesser() *BaikeSearchProcesser {
	return &BaikeSearchProcesser{}
}

func (this *BaiduSearchProcesser) Finish() {
	return
}

func (this *BaikeSearchProcesser) Finish() {
	return
}

var (
	spBaidu = spider.NewSpider(NewBaiduSearchProcesser(), "BaiduSearch")
	spBaike = spider.NewSpider(NewBaikeSearchProcesser(), "BaikeSearch")
)

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

func (this *BaikeSearchProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()
	
	abstract := query.Find(".lemma-summary .para").Text()
	abstract = strings.Trim(abstract, " \t\n\r")
	// the entity we want to save by Pipeline
	p.AddField("answer", abstract)
}


func SearchEngineAnswer(question string) (string, error) {
	baiduAnswer := make(chan string, 1)
	baikeAnswer := make(chan string, 1)

	go baiduSearch(question, baiduAnswer)
	go baikeSearch(question, baikeAnswer)
	
	answer := <- baiduAnswer
	if answer != "" {
		return answer, nil
	}
	answer = <- baikeAnswer
	if answer != "" {
		return answer, nil
	}
	return "", errors.New("empty answer")
}

func baiduSearch(question string, baiduAnswer chan string) {
	
	// GetWithParams Params:
	// 1. url
	// 2. Response type is "html" or "json" or "jsonp" or "text"
	// 3. The urltag is name for marking url and distinguish different urls in PageProcesser and Pipeline.
	// 4. The method is POST or GET
	// 5. The postdata is body string sent to server
	// 6. The header is header for http request.
	// 7. Cookies
	urlFinal := urlBaidu + "/s?" + "wd=" + question
	// urlFinal = "http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn"
	req := request.NewRequest(urlFinal, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := spBaidu.GetByRequest(req)

	baiduAnswer <- pageItems.GetAll()["answer"]
}

func baikeSearch(question string, baikeAnswer chan string) {
	urlFinal := urlBaike + "/item/" + question

	req := request.NewRequest(urlFinal, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := spBaike.GetByRequest(req)
	baikeAnswer <- pageItems.GetAll()["answer"]
}