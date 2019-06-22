package geek

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/pkg/errors"
)

func DeletePic(use bool, article string) string {
	if use && article[len(article)-114:len(article)-107] == "<p><img" {
		return article[:len(article)-114]
	}
	return article
}

func SaveArticleAsPdf(article string, path string) (err error) {
	if article == "" {
		return errors.New("article is empty")
	}
	generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return errors.Wrap(err, "save pdf failed")
	}
	generator.Dpi.Set(300)
	generator.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	html := head + article + tail

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(html)))
	generator.AddPage(page)
	err = generator.Create()
	if err != nil {
		return errors.Wrap(err, "save pdf failed")
	}

	err = generator.WriteFile(path)
	if err != nil {
		return errors.Wrap(err, "save pdf failed")
	}
	return
}

const head = `
<html>
<meta http-equiv="content-type" content="text/html;charset=utf-8">
 <head>
     <style type="text/css">
         ._1kh1ihh6_0 {
    font-weight: 400;
    color: #4c4c4c;
    line-height: 1.76;
    white-space: normal;
    word-break: normal;
    font-size: 16px;
    -webkit-font-smoothing: antialiased;
    -webkit-transition: background-color .3s ease;
    transition: background-color .3s ease;
}
._1kh1ihh6_0 pre {
    margin-top: 16px;
    padding: 35px 0 30px;
    margin-bottom: 30px;
    border-radius: 6px;
    background: rgba(246,247,251,.749);
    border: 0;
}
._1kh1ihh6_0 pre code {
    font-size: 12px;
    font-family: Consolas,Liberation Mono,Menlo,monospace,Courier;
    display: block;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    margin-left: 16px;
    margin-right: 16px;
    overflow: hidden;
    position: relative;
}
._1kh1ihh6_0 img {
    display: block;
    max-width: 100%;
    position: relative;
    left: 50%;
    -webkit-transform: translateX(-50%);
    transform: translateX(-50%);
    background-color: #eee;
    vertical-align: top;
    border-radius: 6px;
}
     </style>
 </head>
 <body>
    <div class="_1kh1ihh6_0">
`
const tail = `</div>
 </body>
</html>
`
