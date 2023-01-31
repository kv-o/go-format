package document

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
