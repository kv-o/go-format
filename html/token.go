package html

import (
	"io"

	"golang.org/x/net/html"
)

var ErrBufferExceeded = html.ErrBufferExceeded

type Attribute = html.Attribute
type Token = html.Token
type TokenType = html.TokenType

const (
	ErrorToken TokenType = iota
	TextToken
	StartTagToken
	EndTagToken
	SelfClosingTagToken
	CommentToken
	DoctypeToken
)

type Tokenizer = html.Tokenizer

func NewTokenizer(r io.Reader) *Tokenizer {
	return html.NewTokenizer(r)
}

func NewTokenizerFragment(r io.Reader, contextTag string) *Tokenizer {
	return html.NewTokenizerFragment(r, contextTag)
}
