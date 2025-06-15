package markdown

import "github.com/thegogod/cmark/ast"

type Markdown struct{}

func (self Markdown) Name() string {
	return "markdown"
}

func (self Markdown) ParseBlock() (ast.Node, error) {
	return nil, nil
}

func (self Markdown) ParseInline() (ast.Node, error) {
	return nil, nil
}
