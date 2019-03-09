package geekpdf

type GeekResp struct {
	Code  int         `json:"code" mapstructure:"code"`
	Data  interface{} `json:"data" mapstructure:"data"`
	Error struct {
		Code int    `json:"code" mapstructure:"code"`
		Msg  string `json:"msg" mapstructure:"msg"`
	} `json:"error" mapstructure:"error"`
	Extra struct {
		Cost      float64 `json:"cost" mapstructure:"cost"`
		RequestID string  `json:"request-id" mapstructure:"request-id"`
	} `json:"extra" mapstructure:"extra"`
}

type LoginReq struct {
	Country   int    `json:"country" mapstructure:"country"`
	Cellphone string `json:"cellphone" mapstructure:"cellphone"`
	Password  string `json:"password" mapstructure:"password"`
	Captcha   string `json:"captcha" mapstructure:"captcha"`
	Remember  int    `json:"remember" mapstructure:"remember"`
	Platform  int    `json:"platform" mapstructure:"platform"`
	Appid     int    `json:"appid" mapstructure:"appid"`
}

type LoginResp struct {
	UID        int    `json:"uid" mapstructure:"uid"`
	Type       int    `json:"type" mapstructure:"type"`
	Cellphone  string `json:"cellphone" mapstructure:"cellphone"`
	Country    string `json:"country" mapstructure:"country"`
	Nickname   string `json:"nickname" mapstructure:"nickname"`
	Avatar     string `json:"avatar" mapstructure:"avatar"`
	Gender     string `json:"gender" mapstructure:"gender"`
	Birthday   string `json:"birthday" mapstructure:"birthday"`
	Graduation string `json:"graduation" mapstructure:"graduation"`
	Profession string `json:"profession" mapstructure:"profession"`
	Industry   string `json:"industry" mapstructure:"industry"`
	Email      string `json:"email" mapstructure:"email"`
	Name       string `json:"name" mapstructure:"name"`
	Address    string `json:"address" mapstructure:"address"`
	Mobile     string `json:"mobile" mapstructure:"mobile"`
	Contact    string `json:"contact" mapstructure:"contact"`
	Position   string `json:"position" mapstructure:"position"`
	Passworded bool   `json:"passworded" mapstructure:"passworded"`
	CreateTime int    `json:"create_time" mapstructure:"create_time"`
	JoinInfoq  string `json:"join_infoq" mapstructure:"join_infoq"`
	OssToken   string `json:"oss_token" mapstructure:"oss_token"`
}

type ArticleReq struct {
	Cid    int    `json:"cid" mapstructure:"cid"`
	Size   int    `json:"size" mapstructure:"size"`
	Prev   int    `json:"prev" mapstructure:"prev"`
	Order  string `json:"order" mapstructure:"order"`
	Sample bool   `json:"sample" mapstructure:"sample"`
}

type ArticleResp struct {
	ArticleSubtitle     string `json:"article_subtitle" mapstructure:"article_subtitle"`
	ID                  int    `json:"id" mapstructure:"id"`
	HadViewed           bool   `json:"had_viewed" mapstructure:"had_viewed"`
	ArticleTitle        string `json:"article_title" mapstructure:"article_title"`
	ArticleCover        string `json:"article_cover" mapstructure:"article_cover"`
	ArticleCouldPreview bool   `json:"article_could_preview" mapstructure:"article_could_preview"`
	ArticleSummary      string `json:"article_summary" mapstructure:"article_summary"`
	ChapterID           string `json:"chapter_id" mapstructure:"chapter_id"`
	Score               int64  `json:"score" mapstructure:"score"`
	ArticleCtime        int    `json:"article_ctime" mapstructure:"article_ctime"`
}
