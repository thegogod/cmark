package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
type UnorderedListElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/ul
func Ul(children ...*ListItemElement) *UnorderedListElement {
	el := Elem("ul")

	for _, child := range children {
		el.Push(child)
	}

	return &UnorderedListElement{el}
}

func (self UnorderedListElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *UnorderedListElement) WithAttr(name string, value string) *UnorderedListElement {
	self.element.WithAttr(name, value)
	return self
}

func (self UnorderedListElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self UnorderedListElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *UnorderedListElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *UnorderedListElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *UnorderedListElement) WithId(value string) *UnorderedListElement {
	self.element.WithId(value)
	return self
}

func (self UnorderedListElement) HasId() bool {
	return self.element.HasId()
}

func (self UnorderedListElement) GetId() string {
	return self.element.GetId()
}

func (self *UnorderedListElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *UnorderedListElement) DelId() {
	self.element.DelId()
}

func (self *UnorderedListElement) WithClass(classes ...string) *UnorderedListElement {
	self.element.WithClass(classes...)
	return self
}

func (self UnorderedListElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self UnorderedListElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *UnorderedListElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *UnorderedListElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *UnorderedListElement) WithStyles(styles ...maps.KeyValue[string, string]) *UnorderedListElement {
	self.element.WithStyles(styles...)
	return self
}

func (self UnorderedListElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *UnorderedListElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *UnorderedListElement) WithStyle(name string, value string) *UnorderedListElement {
	self.element.WithStyle(name, value)
	return self
}

func (self UnorderedListElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self UnorderedListElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *UnorderedListElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *UnorderedListElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self UnorderedListElement) Count() int {
	return self.element.Count()
}

func (self UnorderedListElement) Children() []Node {
	return self.element.Children()
}

func (self *UnorderedListElement) Push(children ...*ListItemElement) *UnorderedListElement {
	for _, child := range children {
		self.element.Push(child)
	}

	return self
}

func (self *UnorderedListElement) Pop() *UnorderedListElement {
	self.element.Pop()
	return self
}

func (self UnorderedListElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self UnorderedListElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *UnorderedListElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *UnorderedListElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
