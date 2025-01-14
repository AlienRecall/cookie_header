package cookie_header

import (
	http "github.com/bogdanfinn/fhttp"
)

type CookieHeader interface {
	LoadFromResponse(res *http.Response, overrideDomain ...string)
	CreateHeader(domains ...string) string
	GetAllCookies(domains ...string) []cookie
	GetCookieValue(name, domain string) string
	AddCookie(name, value, domain string)
	DeleteCookie(name string, domains ...string)
}

// interface check
var _ CookieHeader = (*cookieHeader)(nil)
