package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/tokens"
)

type Markdown struct{}

func New() Markdown {
	return Markdown{}
}

func (self Markdown) Name() string {
	return "markdown"
}

func (self Markdown) ParseBlock(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return nil, nil
}

func (self Markdown) ParseInline(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return nil, nil
}
