package html

import (
	"golang.org/x/net/html"
)

type Node = html.Node
type NodeType = html.NodeType

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
	RawNode
)
