package cmark

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/tokens"
)

type Extension interface {
	Name() string
	ParseBlock(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error)
	ParseInline(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error)
	ParseSyntax(parser ast.Parser, ptr *tokens.Pointer, name string) (ast.Node, error)
}
