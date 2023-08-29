// Package text implements functions for processing text files.
//
// This package provides functions for interoperability with other document file
// formats. This package does not provide functions for general text (file)
// processing, which are provided by the Go standard library.
package text

import (
	"io"

	"git.sr.ht/~kvo/format/document"
)

// Parse returns a pointer to a document.Node representing the contents of the
// given reader.
func Parse(r io.Reader) (*document.Node, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	node := document.Node{
		Type: document.TextNode,
		Data: string(bytes),
	}
	return &node, nil
}
