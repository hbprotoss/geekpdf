# 极客时间导出PDF工具

已购买课程导出PDF工具

## 用法

需要输入极客时间官网账号密码

### wkhtmltopdf

安装请参照官网: [https://wkhtmltopdf.org](https://wkhtmltopdf.org)

### go依赖

```shell
go get github.com/sirupsen/logrus
go get github.com/SebastiaanKlippert/go-wkhtmltopdf
go get github.com/pkg/errors
```

### 参数

```shell
$ ./geekpdf -h
Usage of ./geekpdf:
  -c string
        Cellphone
  -i int
        Product ID
  -p string
        Path to store pdf (default "pdf/")
  -t int
        Request Time Sleep (default 5)
  -w string
        Password
```
