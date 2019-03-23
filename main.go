package main

import (
	"flag"
	"geekpdf/geek"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	articleDownloaderCount = 10
	audioDownloaderCount   = 10
)

var (
	cellphone string
	password  string
	path      string
	cid       int

	articleChan      chan *geek.ArticleListResp
	articleWaitGroup sync.WaitGroup
	audioChan        chan *geek.ArticleResp
	audioWaitGroup   sync.WaitGroup
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

	// init handling chan
	initChan(g)

	for _, articleItem := range articleList {
		articleWaitGroup.Add(1)
		articleChan <- articleItem
	}
	articleWaitGroup.Wait()
	audioWaitGroup.Wait()
}

func downloadArticle(g *geek.GeekTime, articles chan *geek.ArticleListResp, wg *sync.WaitGroup) {
	articleItem := <-articles
	defer wg.Done()

	article, err := g.Article(articleItem.ID)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"articleId": article.ID,
			"title":     article.ArticleTitle,
		}).Error("Loading article failed")
		return
	}
	log.WithFields(log.Fields{
		"articleId": article.ID,
		"title":     article.ArticleTitle,
	}).Info("Loading article success")

	// download audio
	audioWaitGroup.Add(1)
	audioChan <- article

	pdfPath := path + article.ArticleTitle + ".pdf"
	err = geek.SaveArticleAsPdf(article.ArticleContent, pdfPath)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"title":    article.ArticleTitle,
			"filePath": pdfPath,
		}).Error("Save pdf failed")
		return
	}
	log.WithFields(log.Fields{
		"articleId": article.ID,
		"title":     article.ArticleTitle,
		"filePath":  pdfPath,
	}).Info("Save pdf success")
}

func downloadAudio(articles chan *geek.ArticleResp, wg *sync.WaitGroup) {
	article := <-articles
	defer wg.Done()

	url := strings.ReplaceAll(article.AudioDownloadURL, "\\/", "/")
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"url": url,
		}).Error("Http request failed")
		return
	}
	defer resp.Body.Close()

	audioPath := path + article.ArticleTitle + ".mp3"
	file, err := os.OpenFile(audioPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"path": audioPath,
		}).Error("Create file failed")
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"path": audioPath,
		}).Error("Write file failed")
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

	if !strings.HasSuffix(path, "/") {
		path += "/"
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

func initChan(g *geek.GeekTime) {
	articleChan = make(chan *geek.ArticleListResp)
	for i := 0; i < articleDownloaderCount; i++ {
		go func() {
			for {
				downloadArticle(g, articleChan, &articleWaitGroup)
			}
		}()
	}

	audioChan = make(chan *geek.ArticleResp)
	for i := 0; i < audioDownloaderCount; i++ {
		go func() {
			for {
				downloadAudio(audioChan, &audioWaitGroup)
			}
		}()
	}

}
