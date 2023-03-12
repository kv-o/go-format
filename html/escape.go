package html

import (
	"golang.org/x/net/html"
)

func EscapeString(s string) string {
	return html.EscapeString(s)
}

func UnescapeString(s string) string {
	return html.UnescapeString(s)
}
