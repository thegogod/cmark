package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tfoot
type TableFootElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/tfoot
func TFoot(children ...*TableRowElement) *TableFootElement {
	self := &TableFootElement{Elem("tfoot")}
	self.Push(children...)
	return self
}

func (self TableFootElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TableFootElement) WithAttr(name string, value string) *TableFootElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TableFootElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TableFootElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TableFootElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TableFootElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TableFootElement) WithId(value string) *TableFootElement {
	self.element.WithId(value)
	return self
}

func (self TableFootElement) HasId() bool {
	return self.element.HasId()
}

func (self TableFootElement) GetId() string {
	return self.element.GetId()
}

func (self *TableFootElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TableFootElement) DelId() {
	self.element.DelId()
}

func (self *TableFootElement) WithClass(classes ...string) *TableFootElement {
	self.element.WithClass(classes...)
	return self
}

func (self TableFootElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TableFootElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TableFootElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TableFootElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TableFootElement) WithStyles(styles ...maps.KeyValue[string, string]) *TableFootElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TableFootElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableFootElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TableFootElement) WithStyle(name string, value string) *TableFootElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TableFootElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TableFootElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TableFootElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TableFootElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TableFootElement) Count() int {
	return self.element.Count()
}

func (self TableFootElement) Children() []Node {
	return self.element.Children()
}

func (self *TableFootElement) Push(children ...*TableRowElement) *TableFootElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *TableFootElement) Pop() *TableFootElement {
	self.element.Pop()
	return self
}

func (self TableFootElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self TableFootElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *TableFootElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TableFootElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
