package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/thead
type TableHeadElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/thead
func THead(children ...*TableRowElement) *TableHeadElement {
	self := &TableHeadElement{Elem("thead")}
	self.Push(children...)
	return self
}

func (self TableHeadElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TableHeadElement) WithAttr(name string, value string) *TableHeadElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TableHeadElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TableHeadElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TableHeadElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TableHeadElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TableHeadElement) WithId(value string) *TableHeadElement {
	self.element.WithId(value)
	return self
}

func (self TableHeadElement) HasId() bool {
	return self.element.HasId()
}

func (self TableHeadElement) GetId() string {
	return self.element.GetId()
}

func (self *TableHeadElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TableHeadElement) DelId() {
	self.element.DelId()
}

func (self *TableHeadElement) WithClass(classes ...string) *TableHeadElement {
	self.element.WithClass(classes...)
	return self
}

func (self TableHeadElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TableHeadElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TableHeadElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TableHeadElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TableHeadElement) WithStyles(styles ...maps.KeyValue[string, string]) *TableHeadElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TableHeadElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableHeadElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TableHeadElement) WithStyle(name string, value string) *TableHeadElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TableHeadElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TableHeadElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TableHeadElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TableHeadElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TableHeadElement) Count() int {
	return self.element.Count()
}

func (self TableHeadElement) Children() []Node {
	return self.element.Children()
}

func (self *TableHeadElement) Push(children ...*TableRowElement) *TableHeadElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableHeadElement) Pop() *TableHeadElement {
	self.element.Pop()
	return self
}

func (self TableHeadElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self TableHeadElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *TableHeadElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TableHeadElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
