package geek

import (
	"context"
	"time"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/page"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func SaveArticleAsPdf(article string, path string) (err error) {
	if article == "" {
		return errors.New("article is empty")
	}

	// 创建HTML内容
	html := head + article + tail

	// 将HTML写入临时文件
	tempFile, err := ioutil.TempFile("", "geek_article_*.html")
	if err != nil {
		return errors.Wrap(err, "create temp html file failed")
	}
	defer os.Remove(tempFile.Name()) // 清理临时文件

	_, err = tempFile.WriteString(html)
	if err != nil {
		return errors.Wrap(err, "write html to temp file failed")
	}
	tempFile.Close()

	// 设置chromedp选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 创建新上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 生成PDF
	var buf []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate("file://"+tempFile.Name()),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 等待页面加载完成
			// 这里可以等待特定元素加载完成，如果需要的话
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 使用默认选项生成PDF
			// 需要使用正确的chromedp API
			var pdfBytes []byte
			var err error
			pdfBytes, _, err = page.PrintToPDF().
				WithTransferMode(page.PrintToPDFTransferModeReturnAsBase64).
				WithScale(1.0).
				WithPaperWidth(8.27).        // A4 width in inches
				WithPaperHeight(11.7).       // A4 height in inches
				WithMarginTop(0.4).
				WithMarginBottom(0.4).
				WithMarginLeft(0.4).
				WithMarginRight(0.4).
				WithPrintBackground(true).
				WithDisplayHeaderFooter(false).
				WithLandscape(false).
				Do(ctx)
			if err != nil {
				return err
			}

			// 解码Base64字符串到字节数组
			buf = pdfBytes
			return nil
		}),
	)
	if err != nil {
		return errors.Wrap(err, "generate pdf failed")
	}

	// 写入PDF文件
	err = ioutil.WriteFile(path, buf, 0644)
	if err != nil {
		return errors.Wrap(err, "write pdf file failed")
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
