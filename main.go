package main

import (
	"geekpdf/geekpdf"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := Init()
	if err != nil {
		log.WithError(err).Error("init failed")
	}

	cellphone := "xx"
	password := "xx"
	path := "/Users/hbprotoss/Downloads/geek/"
	cid := 139

	g := geekpdf.NewGeekTime(cellphone, password)

	loginResp, err := g.Login()
	if err != nil {
		log.WithError(err).Error("Login failed")
		return
	}
	log.WithFields(log.Fields{
		"uid":     loginResp.UID,
		"cookies": g.Cookies(),
	}).Info("Login success")

	articleList, err := g.ArticleList(cid)
	if err != nil {
		log.WithError(err).Error("Loading article list failed")
		return
	}
	log.WithFields(log.Fields{
		"productId": cid,
		"count":     len(articleList),
	}).Info("Loading article list success")

	for _, article := range articleList {
		article, err := g.Article(article.ID)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"articleId": article.ID,
				"title":     article.ArticleTitle,
			}).Error("Loading article failed")
			continue
		}
		log.WithFields(log.Fields{
			"articleId": article.ID,
			"title":     article.ArticleTitle,
		}).Info("Loading article success")

		pdfPath := path + article.ArticleTitle + ".pdf"
		err = geekpdf.SaveArticleAsPdf(article.ArticleContent, pdfPath)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"title":    article.ArticleTitle,
				"filePath": pdfPath,
			}).Error("Save pdf failed")
			continue
		}
		log.WithFields(log.Fields{
			"articleId": article.ID,
			"title":     article.ArticleTitle,
			"filePath":  pdfPath,
		}).Info("Save pdf success")
	}
}

func Init() (err error) {
	initLog()
	return nil
}

func initLog() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}
