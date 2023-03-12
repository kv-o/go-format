// Package html implements an HTML5-compliant tokenizer and parser.
//
// Currently, this package serves as a wrapper for the Go development team's
// official HTML5 processing library ("golang.org/x/net/html"). For library
// documentation, consult the upstream Go library.
//
// This package will most likely be rewritten in the future to use document.Node
// and other structures common to all document formats instead of html.Node,
// which is HTML-specific.
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
