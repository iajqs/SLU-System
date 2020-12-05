/**
* Created by cks
* Date: 2020-12-04
* Time: 09:47
*/
package site

import (
	"SLU-System/config"
	"github.com/sirupsen/logrus"
	// "fmt"
	"net/http"
)

type Site struct {
}

func New() *Site {
	return &Site{}
}

func (s *Site) Run() {
	siteConfig := config.Conf.Site
	port := siteConfig.SiteBase.ListenPort
	//bind := siteConfig.SiteBase.Bind
	addr := fmt.Sprintf(":%d", port)
	logrus.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("./site/"))))
}
