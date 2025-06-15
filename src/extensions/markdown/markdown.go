package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/tokens"
)

type Markdown struct{}

func (self Markdown) Name() string {
	return "markdown"
}

func (self Markdown) ParseBlock(ptr *tokens.Pointer) (ast.Node, error) {
	return nil, nil
}

func (self Markdown) ParseInline(ptr *tokens.Pointer) (ast.Node, error) {
	return nil, nil
}
