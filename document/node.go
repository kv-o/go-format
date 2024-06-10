// Package document provides identifiers and structures common to all document
// file formats.
package document

import (
	"git.sr.ht/~kvo/go-format/html"
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
