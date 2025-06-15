package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ol
type OrderedListElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ol
func Ol(children ...*ListItemElement) *OrderedListElement {
	el := Elem("ol")

	for _, child := range children {
		el.Push(child)
	}

	return &OrderedListElement{el}
}

func (self OrderedListElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *OrderedListElement) WithAttr(name string, value string) *OrderedListElement {
	self.element.WithAttr(name, value)
	return self
}

func (self OrderedListElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self OrderedListElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *OrderedListElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *OrderedListElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *OrderedListElement) WithId(value string) *OrderedListElement {
	self.element.WithId(value)
	return self
}

func (self OrderedListElement) HasId() bool {
	return self.element.HasId()
}

func (self OrderedListElement) GetId() string {
	return self.element.GetId()
}

func (self *OrderedListElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *OrderedListElement) DelId() {
	self.element.DelId()
}

func (self *OrderedListElement) WithClass(classes ...string) *OrderedListElement {
	self.element.WithClass(classes...)
	return self
}

func (self OrderedListElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self OrderedListElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *OrderedListElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *OrderedListElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *OrderedListElement) WithStyles(styles ...maps.KeyValue[string, string]) *OrderedListElement {
	self.element.WithStyles(styles...)
	return self
}

func (self OrderedListElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *OrderedListElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *OrderedListElement) WithStyle(name string, value string) *OrderedListElement {
	self.element.WithStyle(name, value)
	return self
}

func (self OrderedListElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self OrderedListElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *OrderedListElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *OrderedListElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self OrderedListElement) Count() int {
	return self.element.Count()
}

func (self OrderedListElement) Children() []Node {
	return self.element.Children()
}

func (self *OrderedListElement) Push(children ...*ListItemElement) *OrderedListElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *OrderedListElement) Pop() *OrderedListElement {
	self.element.Pop()
	return self
}

func (self OrderedListElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self OrderedListElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *OrderedListElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *OrderedListElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
