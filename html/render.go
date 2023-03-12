package html

import (
	"io"

	"golang.org/x/net/html"
)

func Render(w io.Writer, n *Node) error {
	return html.Render(w, n)
}
