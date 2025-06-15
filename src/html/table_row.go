package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tr
type TableRowElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tr
func Tr(children ...*TableCellElement) *TableRowElement {
	self := &TableRowElement{Elem("tr")}
	self.Push(children...)
	return self
}

func (self TableRowElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TableRowElement) WithAttr(name string, value string) *TableRowElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TableRowElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TableRowElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TableRowElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TableRowElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TableRowElement) WithId(value string) *TableRowElement {
	self.element.WithId(value)
	return self
}

func (self TableRowElement) HasId() bool {
	return self.element.HasId()
}

func (self TableRowElement) GetId() string {
	return self.element.GetId()
}

func (self *TableRowElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TableRowElement) DelId() {
	self.element.DelId()
}

func (self *TableRowElement) WithClass(classes ...string) *TableRowElement {
	self.element.WithClass(classes...)
	return self
}

func (self TableRowElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TableRowElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TableRowElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TableRowElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TableRowElement) WithStyles(styles ...maps.KeyValue[string, string]) *TableRowElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TableRowElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableRowElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TableRowElement) WithStyle(name string, value string) *TableRowElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TableRowElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TableRowElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TableRowElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TableRowElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TableRowElement) Count() int {
	return self.element.Count()
}

func (self TableRowElement) Children() []Node {
	return self.element.Children()
}

func (self *TableRowElement) Push(children ...*TableCellElement) *TableRowElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableRowElement) Pop() *TableRowElement {
	self.element.Pop()
	return self
}

func (self TableRowElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self TableRowElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *TableRowElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TableRowElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
