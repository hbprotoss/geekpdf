# 极客时间导出PDF工具

已购买课程导出PDF工具

## 用法

需要输入极客时间官网账号密码

本工具现在使用chromedp库进行HTML到PDF转换，基于Chrome/Chromium浏览器引擎，已不再依赖wkhtmltopdf。

### go依赖

依赖管理通过Go modules自动处理，运行以下命令即可安装依赖：

```shell
go mod tidy
```

### 参数

```shell
$ ./geekpdf -h
Usage of ./geekpdf:
  -c string
    	Cellphone
  -f string
    	Cookie file (optional, bypasses login if provided)
  -i int
    	Product ID
  -p string
    	Path to store pdf (default "pdf/")
  -w string
    	Password
```