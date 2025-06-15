package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
type ListItemElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/li
func Li(children ...any) *ListItemElement {
	return &ListItemElement{Elem("li").Push(children...)}
}

func (self ListItemElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *ListItemElement) WithAttr(name string, value string) *ListItemElement {
	self.element.WithAttr(name, value)
	return self
}

func (self ListItemElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self ListItemElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *ListItemElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *ListItemElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *ListItemElement) WithId(value string) *ListItemElement {
	self.element.WithId(value)
	return self
}

func (self ListItemElement) HasId() bool {
	return self.element.HasId()
}

func (self ListItemElement) GetId() string {
	return self.element.GetId()
}

func (self *ListItemElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *ListItemElement) DelId() {
	self.element.DelId()
}

func (self *ListItemElement) WithClass(classes ...string) *ListItemElement {
	self.element.WithClass(classes...)
	return self
}

func (self ListItemElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self ListItemElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *ListItemElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *ListItemElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *ListItemElement) WithStyles(styles ...maps.KeyValue[string, string]) *ListItemElement {
	self.element.WithStyles(styles...)
	return self
}

func (self ListItemElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ListItemElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *ListItemElement) WithStyle(name string, value string) *ListItemElement {
	self.element.WithStyle(name, value)
	return self
}

func (self ListItemElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self ListItemElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *ListItemElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *ListItemElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self ListItemElement) Count() int {
	return self.element.Count()
}

func (self ListItemElement) Children() []Node {
	return self.element.Children()
}

func (self *ListItemElement) Push(children ...any) *ListItemElement {
	self.element.Push(children...)
	return self
}

func (self *ListItemElement) Pop() *ListItemElement {
	self.element.Pop()
	return self
}

func (self ListItemElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self ListItemElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *ListItemElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *ListItemElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
