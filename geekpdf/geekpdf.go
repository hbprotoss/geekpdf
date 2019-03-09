package geekpdf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type GeekTime struct {
	Cellphone string
	Password  string

	client  *http.Client
	cookies []*http.Cookie
}

func (g *GeekTime) Cookies() []*http.Cookie {
	return g.cookies
}

func NewGeekTime(cellphone string, password string) *GeekTime {
	client := &http.Client{}

	return &GeekTime{
		Cellphone: cellphone,
		Password:  password,

		client: client,
	}
}

func (g *GeekTime) Login() (loginResp *LoginResp, err error) {
	url := "https://account.geekbang.org/account/ticket/login"
	body := &LoginReq{
		Cellphone: g.Cellphone,
		Password:  g.Password,
		Country:   86,
		Remember:  1,
		Platform:  3,
		Appid:     1,
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

func (g *GeekTime) Articles(cid int) (err error) {
	url := "https://time.geekbang.org/serv/v1/column/articles"
	body := &ArticleReq{
		Cid:    cid,
		Size:   100,
		Prev:   0,
		Order:  "earliest",
		Sample: false,
	}
	response, err := g.post(url, body)
	fmt.Printf("%v", response)
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
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal error")
	}

	request, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")

	if g.cookies != nil {
		for _, cookie := range g.cookies {
			request.AddCookie(cookie)
		}
	}
	return
}

func (g *GeekTime) parseResponse(resp *http.Response, output interface{}) (err error) {
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
		return errors.New(geek.Error.Msg)
	}
	return mapstructure.Decode(geek.Data, output)
}
