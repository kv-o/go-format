package markdown

import (
	"errors"

	"golang.org/x/net/html"
)

type NodeType uint32

const (
	BoldNode NodeType = iota
	BreakNode
	CodeNode
	H1Node
	H2Node
	H3Node
	H4Node
	H5Node
	H6Node
	ImgNode
	InCodeNode
	ItalicNode
	LineNode
	LinkNode
	OlNode
	ParaNode
	UlNode
	QuoteNode
)

type Node struct {
	Type NodeType
	Data string
	Attr map[string]string
}

func parseHtml(r *html.Node) (*Node, error) {
	// TODO: Implement this function.
	return nil, errors.New("unimplemented function")
}

// Parse returns the parse tree for a document contained in interface r, which
// at the moment may only be an html.Node.
func Parse(r any) (*Node, error) {
	switch r.(type) {
	case html.Node:
		return parseHtml(r)
	default:
		return nil, errors.New("invalid parse input type")
	}
}
