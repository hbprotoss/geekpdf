package geekpdf

var (
	anonUrlSet = map[string]bool{
		"https://account.geekbang.org/account/ticket/login": true,
	}
)

func needLogin(url string) bool {
	need, ok := anonUrlSet[url]
	return !ok || !need
}
