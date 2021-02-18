/**
* Created by cks
* Date: 2020-11-26
* Time: 20:06
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
	urlBaike = "http://baike.baidu.com"
	// userAgent = "Mozilla/5.0 BaiduSpider/2.0"
)

type BaikeSearchProcesser struct {
}

func NewBaikeSearchProcesser() *BaikeSearchProcesser {
	return &BaikeSearchProcesser{}
}

func (this *BaikeSearchProcesser) Finish() {
	return
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


func BaikeAnswer(question string) (string, error) {
	answer := baikeSearch(question)
	
	if answer != "" {
		return answer, nil
	}
	return "", errors.New("empty answer")
}

func baikeSearch(question string) string {
	spBaike := spider.NewSpider(NewBaikeSearchProcesser(), "BaikeSearch")
	urlFinal := urlBaike + "/item/" + question

	req := request.NewRequest(urlFinal, "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := spBaike.GetByRequest(req)
	if pageItems == nil {
		return ""
	}
	return pageItems.GetAll()["answer"]
}