package main

import (
	"fmt"
	"geekpdf/geekpdf"
)

func main() {
	g := geekpdf.NewGeekTime("xxx", "xxx")
	resp, err := g.Login()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	cookies := g.Cookies()
	fmt.Printf("%v\n", cookies)
	fmt.Printf("%v\n", resp)
}
