package html

import (
	"io"

	"golang.org/x/net/html"
)

func Parse(r io.Reader) (*Node, error) {
	return html.Parse(r)
}

func ParseFragment(r io.Reader, context *Node) ([]*Node, error) {
	return html.ParseFragment(r, context)
}

func ParseFragmentWithOptions(r io.Reader, context *Node, opts ...ParseOption) ([]*Node, error) {
	return html.ParseFragmentWithOptions(r, context, opts...)
}

func ParseWithOptions(r io.Reader, opts ...ParseOption) (*Node, error) {
	return html.ParseWithOptions(r, opts...)
}

type ParseOption = html.ParseOption

func ParseOptionEnableScripting(enable bool) ParseOption {
	return html.ParseOptionEnableScripting(enable)
}
