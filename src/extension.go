package cmark

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/tokens"
)

type Extension interface {
	Name() string
	ParseBlock(ptr *tokens.Pointer) (ast.Node, error)
	ParseInline(ptr *tokens.Pointer) (ast.Node, error)
}
