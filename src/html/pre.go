package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
type PreElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/pre
func Pre(children ...any) *PreElement {
	return &PreElement{Elem("pre").Push(children...)}
}

func (self PreElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *PreElement) WithAttr(name string, value string) *PreElement {
	self.element.WithAttr(name, value)
	return self
}

func (self PreElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self PreElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *PreElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *PreElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *PreElement) WithId(value string) *PreElement {
	self.element.WithId(value)
	return self
}

func (self PreElement) HasId() bool {
	return self.element.HasId()
}

func (self PreElement) GetId() string {
	return self.element.GetId()
}

func (self *PreElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *PreElement) DelId() {
	self.element.DelId()
}

func (self *PreElement) WithClass(classes ...string) *PreElement {
	self.element.WithClass(classes...)
	return self
}

func (self PreElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self PreElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *PreElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *PreElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *PreElement) WithStyles(styles ...maps.KeyValue[string, string]) *PreElement {
	self.element.WithStyles(styles...)
	return self
}

func (self PreElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *PreElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *PreElement) WithStyle(name string, value string) *PreElement {
	self.element.WithStyle(name, value)
	return self
}

func (self PreElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self PreElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *PreElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *PreElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self PreElement) Count() int {
	return self.element.Count()
}

func (self PreElement) Children() []Node {
	return self.element.Children()
}

func (self *PreElement) Push(children ...any) *PreElement {
	self.element.Push(children...)
	return self
}

func (self *PreElement) Pop() *PreElement {
	self.element.Pop()
	return self
}

func (self PreElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self PreElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *PreElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *PreElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
