package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

type AstElement []ast.Node

func Ast(nodes ...ast.Node) *AstElement {
	self := AstElement{}
	self = append(self, nodes...)
	return &self
}

func (self AstElement) GetSelector() string {
	return ""
}

func (self AstElement) HasAttr(name string) bool {
	return false
}

func (self AstElement) GetAttr(name string) string {
	return ""
}

func (self AstElement) SetAttr(name string, value string) {
	return
}

func (self AstElement) DelAttr(name string) {
	return
}

func (self AstElement) HasId() bool {
	return false
}

func (self AstElement) GetId() string {
	return ""
}

func (self AstElement) SetId(id string) {
	return
}

func (self AstElement) DelId() {
	return
}

func (self AstElement) HasClass(name ...string) bool {
	return false
}

func (self AstElement) GetClass() []string {
	return []string{}
}

func (self AstElement) AddClass(name ...string) {
	return
}

func (self AstElement) DelClass(name ...string) {
	return
}

func (self AstElement) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self AstElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self AstElement) HasStyle(name ...string) bool {
	return false
}

func (self AstElement) GetStyle(name string) string {
	return ""
}

func (self AstElement) SetStyle(name string, value string) {
	return
}

func (self AstElement) DelStyle(name ...string) {
	return
}

func (self AstElement) Count() int {
	return len(self)
}

func (self AstElement) Validate(scope *ast.Scope) error {
	child := scope.Create()

	for _, node := range self {
		if err := node.Validate(child); err != nil {
			return err
		}
	}

	return nil
}

func (self AstElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	child := scope.Create()

	for _, node := range self {
		value, err := node.Evaluate(child)

		if err != nil {
			return value, err
		}

		if !value.IsNil() {
			return value, nil
		}
	}

	return reflect.NewNil(), nil
}

func (self AstElement) Render(scope *ast.Scope) []byte {
	child := scope.Create()
	value := []byte{}

	for _, node := range self {
		block, err := node.Evaluate(child)

		if err != nil {
			continue
		}

		value = append(value, block.String()...)
	}

	return value
}

func (self AstElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	child := scope.Create()
	value := []byte{}

	for _, node := range self {
		block, err := node.Evaluate(child)

		if err != nil {
			continue
		}

		value = append(value, block.String()...)
	}

	return value
}

func (self AstElement) GetById(id string) Node {
	return nil
}

func (self AstElement) Select(query ...any) []Node {
	return []Node{}
}
