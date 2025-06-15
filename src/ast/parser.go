package ast

import (
	"github.com/thegogod/cmark/tokens"
)

type Parser interface {
	Parse(src []byte) (Node, error)
	ParseBlock(ptr *tokens.Pointer) (Node, error)
	ParseInline(ptr *tokens.Pointer) (Node, error)
}
