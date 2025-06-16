package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tbody
type TableBodyElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tbody
func TBody(children ...*TableRowElement) *TableBodyElement {
	self := &TableBodyElement{Elem("tbody")}
	self.Push(children...)
	return self
}

func (self TableBodyElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TableBodyElement) WithAttr(name string, value string) *TableBodyElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TableBodyElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TableBodyElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TableBodyElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TableBodyElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TableBodyElement) WithId(value string) *TableBodyElement {
	self.element.WithId(value)
	return self
}

func (self TableBodyElement) HasId() bool {
	return self.element.HasId()
}

func (self TableBodyElement) GetId() string {
	return self.element.GetId()
}

func (self *TableBodyElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TableBodyElement) DelId() {
	self.element.DelId()
}

func (self *TableBodyElement) WithClass(classes ...string) *TableBodyElement {
	self.element.WithClass(classes...)
	return self
}

func (self TableBodyElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TableBodyElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TableBodyElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TableBodyElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TableBodyElement) WithStyles(styles ...maps.KeyValue[string, string]) *TableBodyElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TableBodyElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableBodyElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TableBodyElement) WithStyle(name string, value string) *TableBodyElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TableBodyElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TableBodyElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TableBodyElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TableBodyElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TableBodyElement) Count() int {
	return self.element.Count()
}

func (self TableBodyElement) Children() []Node {
	return self.element.Children()
}

func (self *TableBodyElement) Push(children ...*TableRowElement) *TableBodyElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableBodyElement) Pop() *TableBodyElement {
	self.element.Pop()
	return self
}

func (self TableBodyElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self TableBodyElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *TableBodyElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TableBodyElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self TableBodyElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self TableBodyElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
