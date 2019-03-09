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
	AppId     int    `json:"appid"`
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

type ArticleListReq struct {
	Cid    int    `json:"cid"`
	Size   int    `json:"size"`
	Prev   int    `json:"prev"`
	Order  string `json:"order"`
	Sample bool   `json:"sample"`
}

type ArticleListResp struct {
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

type ArticleReq struct {
	ID string `json:"id"`
}

type ArticleResp struct {
	Sku                 string `json:"sku"`
	VideoCover          string `json:"video_cover"`
	AuthorName          string `json:"author_name"`
	TextReadVersion     int    `json:"text_read_version"`
	ArticleCover        string `json:"article_cover"`
	ProductType         string `json:"product_type"`
	AudioURL            string `json:"audio_url"`
	ChapterID           string `json:"chapter_id"`
	ColumnHadSub        bool   `json:"column_had_sub"`
	AudioDubber         string `json:"audio_dubber"`
	AudioTime           string `json:"audio_time"`
	VideoHeight         int    `json:"video_height"`
	ArticleContent      string `json:"article_content"`
	ArticleCoverHidden  bool   `json:"article_cover_hidden"`
	ColumnIsExperience  bool   `json:"column_is_experience"`
	Score               string `json:"score"`
	VideoMedia          string `json:"video_media"`
	ArticleSubtitle     string `json:"article_subtitle"`
	AudioDownloadURL    string `json:"audio_download_url"`
	ID                  int    `json:"id"`
	HadViewed           bool   `json:"had_viewed"`
	ArticleTitle        string `json:"article_title"`
	ColumnBgColor       string `json:"column_bgcolor"`
	ArticleSummary      string `json:"article_summary"`
	VideoTime           string `json:"video_time"`
	ProductID           int    `json:"product_id"`
	ArticlePosterWxLite string `json:"article_poster_wxlite"`
	LikeCount           int    `json:"like_count"`
	HadLiked            bool   `json:"had_liked"`
	ColumnID            int    `json:"column_id"`
	ColumnCover         string `json:"column_cover"`
	AudioTimeArr        struct {
		M string `json:"m"`
		S string `json:"s"`
		H string `json:"h"`
	} `json:"audio_time_arr"`
	AudioTitle          string `json:"audio_title"`
	AudioSize           int    `json:"audio_size"`
	AudioMd5            string `json:"audio_md5"`
	TextReadPercent     int    `json:"text_read_percent"`
	ArticleShareTitle   string `json:"article_sharetitle"`
	Cid                 int    `json:"cid"`
	VideoSize           int    `json:"video_size"`
	ViewCount           int    `json:"view_count"`
	VideoWidth          int    `json:"video_width"`
	ColumnCouldSub      bool   `json:"column_could_sub"`
	ArticleCtime        int    `json:"article_ctime"`
	ArticleCouldPreview bool   `json:"article_could_preview"`
}
