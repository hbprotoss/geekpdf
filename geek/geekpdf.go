package geek

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))

// getRandomUserAgent returns a random user agent string from available browsers with randomized version numbers
func getRandomUserAgent() string {
	// Generate random major versions (80+ to ensure 2020+ compatibility)
	chromeVersion := randomSource.Intn(20) + 100  // Generates a number between 100-119
	firefoxVersion := randomSource.Intn(25) + 115 // Generates a number between 115-139
	edgeVersion := randomSource.Intn(20) + 100    // Generates a number between 100-119

	// Format user agents with random versions
	userAgents := []string{
		fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36", chromeVersion),
		fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:%d.0) Gecko/20100101 Firefox/%d.0", firefoxVersion, firefoxVersion),
		fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.0.0 Safari/537.36 Edge/%d.0.0.0", chromeVersion, edgeVersion),
	}
	return userAgents[randomSource.Intn(len(userAgents))]
}

type GeekTime struct {
	Cellphone string
	Password  string

	client  *http.Client
	cookies []*http.Cookie
}

func (g *GeekTime) Cookies() []*http.Cookie {
	return g.cookies
}

func NewGeekTime(cellphone string, password string, cookieFile string) *GeekTime {
	client := &http.Client{}
	var cookies []*http.Cookie

	if cookieFile != "" {
		var err error
		cookies, err = loadCookiesFromFile(cookieFile)
		if err != nil {
			panic("load cookies from file failed: " + err.Error())
		}
	}

	return &GeekTime{
		Cellphone: cellphone,
		Password:  password,
		client:    client,
		cookies:   cookies,
	}
}

func loadCookiesFromFile(filePath string) ([]*http.Cookie, error) {
	// Read the cookie file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read cookie file")
	}

	// Convert to string and trim whitespace
	cookieStr := strings.TrimSpace(string(data))
	if cookieStr == "" {
		return nil, errors.New("cookie file is empty")
	}

	var cookies []*http.Cookie

	// Split by semicolon to get individual cookie pairs
	cookiePairs := strings.Split(cookieStr, ";")

	for _, pair := range cookiePairs {
		// Trim whitespace from each pair
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		// Split by the first equals sign to get name and value
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			// Skip invalid cookie pairs
			continue
		}

		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if name == "" {
			continue
		}

		// Create http.Cookie
		cookie := &http.Cookie{
			Name:  name,
			Value: value,
		}
		cookies = append(cookies, cookie)
	}

	if len(cookies) == 0 {
		return nil, errors.New("no valid cookies found in file")
	}

	return cookies, nil
}

func (g *GeekTime) Login() (loginResp *LoginResp, err error) {
	if len(g.cookies) > 0 {
		// already have cookies, skip login
		return &LoginResp{
			UID: 0,
		}, nil
	}

	url := "https://account.geekbang.org/account/ticket/login"
	body := &LoginReq{
		Cellphone: g.Cellphone,
		Password:  g.Password,
		Country:   86,
		Remember:  1,
		Platform:  4,
		AppId:     1,
	}
	header := &http.Header{}
	header.Add("Referer", "https://account.geekbang.org/login?redirect=https%3A%2F%2Ftime.geekbang.org%2F")

	resp, err := g.postWithHeader(url, body, header)
	if err != nil {
		return nil, errors.Wrap(err, "login failed")
	}
	g.cookies = resp.Cookies()

	loginResp = &LoginResp{}
	err = g.parseResponse(resp, loginResp)
	if err != nil {
		return nil, err
	}
	return
}

func (g *GeekTime) ArticleList(cid int) (articles []*ArticleListResp, err error) {
	url := "https://time.geekbang.org/serv/v1/column/articles"
	body := &ArticleListReq{
		Cid:    cid,
		Size:   100,
		Prev:   0,
		Order:  "earliest",
		Sample: false,
	}
	response, err := g.post(url, body)
	if err != nil {
		return nil, errors.New("get article list failed")
	}

	data := GeekList{}
	err = g.parseResponse(response, &data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data.List, &articles)
	return
}

func (g *GeekTime) Article(id int) (article *ArticleResp, err error) {
	url := "https://time.geekbang.org/serv/v1/article"
	body := &ArticleReq{
		ID: strconv.Itoa(id),
	}
	response, err := g.post(url, body)
	if err != nil {
		return nil, errors.New("get article failed")
	}

	article = &ArticleResp{}
	err = g.parseResponse(response, article)
	if err != nil {
		return nil, err
	}
	return
}

func (g *GeekTime) post(url string, body interface{}) (response *http.Response, err error) {
	request, err := g.makeRequest(url, body)
	if err != nil {
		return nil, errors.Wrap(err, "makeRequest error")
	}
	response, err = g.client.Do(request)
	return
}

func (g *GeekTime) postWithHeader(url string, body interface{}, header *http.Header) (response *http.Response, err error) {
	request, err := g.makeRequest(url, body)
	if err != nil {
		return nil, errors.Wrap(err, "makeRequest error")
	}
	for k, v := range *header {
		request.Header[k] = v
	}
	response, err = g.client.Do(request)
	return
}

func (g *GeekTime) makeRequest(url string, body interface{}) (request *http.Request, err error) {
	if needLogin(url) && g.cookies == nil {
		return nil, errors.New("need login")
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal error")
	}

	request, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Origin", "https://time.geekbang.org")
	request.Header.Add("User-Agent", getRandomUserAgent())

	for _, cookie := range g.cookies {
		request.AddCookie(cookie)
	}
	return
}

func (g *GeekTime) parseResponse(resp *http.Response, output interface{}) (err error) {
	if resp.StatusCode != http.StatusOK {
		return errors.New("http code " + resp.Status)
	}

	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read body error")
	}

	geek := &GeekResp{}
	err = json.Unmarshal(raw, geek)
	if err != nil {
		return errors.Wrap(err, "parse outer response error")
	}

	if geek.Code != 0 {
		gerr := &GeekError{}
		err = json.Unmarshal(geek.Error, gerr)
		if err != nil {
			return errors.Wrap(err, "parse error msg failed")
		}
		return errors.New(gerr.Msg)
	}
	return json.Unmarshal(geek.Data, output)
}
