package markdown

import (
	"errors"
)

func parseHtml(r *html.Node) (string, error) {
	// TODO: Implement this function.
	return "", errors.New("unimplemented function")
}

// Parse returns the parse tree for a document contained in interface r, which
// at the moment may only be an html.Node.
func Parse(r any) (string, error) {
	switch r.(type) {
	case html.Node:
		return parseHtml(r)
	default:
		return nil, errors.New("invalid parse input type")
	}
}
