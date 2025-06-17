package cmark

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

type Extension interface {
	Name() string
	ParseBlock(parser html.Parser, ptr *tokens.Pointer) (html.Node, error)
	ParseInline(parser html.Parser, ptr *tokens.Pointer) (html.Node, error)
}
