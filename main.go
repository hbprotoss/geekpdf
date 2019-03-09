package main

import (
	"flag"
	"geekpdf/geek"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	cellphone string
	password  string
	path      string
	cid       int
)

func main() {
	err := Init()
	if err != nil {
		log.WithError(err).Error("init failed")
	}

	g := geek.NewGeekTime(cellphone, password)

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
		err = geek.SaveArticleAsPdf(article.ArticleContent, pdfPath)
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
	initCmd()
	initApp()
	return nil
}

func initLog() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

func initCmd() {
	flag.StringVar(&cellphone, "c", "", "Cellphone")
	flag.StringVar(&password, "w", "", "Password")
	flag.StringVar(&path, "p", "pdf/", "Path to store pdf")
	flag.IntVar(&cid, "i", 0, "Product ID")
	flag.Parse()

	if cellphone == "" || password == "" {
		log.Error("Invalid cellphone or password")
		os.Exit(1)
	}
}

func initApp() {
	if _, err := os.Stat(path); err == nil {
		return
	}
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.WithError(err).Error("Can not init dir for saving pdf")
		os.Exit(1)
	}
}
