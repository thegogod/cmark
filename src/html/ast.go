package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

type AstElement []ast.Node

func Ast(nodes ...ast.Node) AstElement {
	return nodes
}

func (self AstElement) HasId() bool {
	return false
}

func (self AstElement) GetId() string {
	return ""
}

func (self *AstElement) SetId(id string) {

}

func (self *AstElement) DelId() {

}

func (self AstElement) HasClass(classes ...string) bool {
	return false
}

func (self AstElement) GetClass() []string {
	return []string{}
}

func (self *AstElement) AddClass(name ...string) {

}

func (self *AstElement) DelClass(name ...string) {

}

func (self AstElement) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self *AstElement) SetStyles(styles ...maps.KeyValue[string, string]) {

}

func (self AstElement) HasStyle(name ...string) bool {
	return false
}

func (self AstElement) GetStyle(name string) string {
	return ""
}

func (self *AstElement) SetStyle(name string, value string) {

}

func (self *AstElement) DelStyle(name ...string) {

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
