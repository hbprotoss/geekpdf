package geekpdf

import "encoding/json"

type GeekResp struct {
	Code  int             `json:"code"`
	Data  json.RawMessage `json:"data"`
	Error json.RawMessage `json:"error"`
}

type GeekError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GeekList struct {
	List json.RawMessage `json:"list"`
	Page struct {
		Count int  `json:"count"`
		More  bool `json:"more"`
	} `json:"page"`
}

type LoginReq struct {
	Country   int    `json:"country"`
	Cellphone string `json:"cellphone"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	Remember  int    `json:"remember"`
	Platform  int    `json:"platform"`
	Appid     int    `json:"appid"`
}

type LoginResp struct {
	UID        int    `json:"uid"`
	Type       int    `json:"type"`
	Cellphone  string `json:"cellphone"`
	Country    string `json:"country"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Graduation string `json:"graduation"`
	Profession string `json:"profession"`
	Industry   string `json:"industry"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Mobile     string `json:"mobile"`
	Contact    string `json:"contact"`
	Position   string `json:"position"`
	Passworded bool   `json:"passworded"`
	CreateTime int    `json:"create_time"`
	JoinInfoq  string `json:"join_infoq"`
	OssToken   string `json:"oss_token"`
}

type ArticleReq struct {
	Cid    int    `json:"cid"`
	Size   int    `json:"size"`
	Prev   int    `json:"prev"`
	Order  string `json:"order"`
	Sample bool   `json:"sample"`
}

type ArticleResp struct {
	ArticleSubtitle     string `json:"article_subtitle"`
	ID                  int    `json:"id"`
	HadViewed           bool   `json:"had_viewed"`
	ArticleTitle        string `json:"article_title"`
	ArticleCover        string `json:"article_cover"`
	ArticleCouldPreview bool   `json:"article_could_preview"`
	ArticleSummary      string `json:"article_summary"`
	ChapterID           string `json:"chapter_id"`
	Score               int64  `json:"score"`
	ArticleCtime        int    `json:"article_ctime"`
}
