package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

type Node interface {
	GetSelector() string

	HasAttr(name string) bool
	GetAttr(name string) string
	SetAttr(name string, value string)
	DelAttr(name string)

	HasId() bool
	GetId() string
	SetId(id string)
	DelId()

	HasClass(name ...string) bool
	GetClass() []string
	AddClass(name ...string)
	DelClass(name ...string)

	GetStyles() maps.OMap[string, string]
	SetStyles(styles ...maps.KeyValue[string, string])

	HasStyle(name ...string) bool
	GetStyle(name string) string
	SetStyle(name string, value string)
	DelStyle(name ...string)

	GetById(id string) Node
	Select(query ...any) []Node

	Render(scope *ast.Scope) []byte
	RenderPretty(scope *ast.Scope, indent string) []byte
}

type ParentNode interface {
	Node

	Count() int
	Children() []Node
}
