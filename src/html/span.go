package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/span
type SpanElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/span
func Span(children ...any) *SpanElement {
	return &SpanElement{Elem("span").Push(children...)}
}

func (self SpanElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *SpanElement) WithAttr(name string, value string) *SpanElement {
	self.element.WithAttr(name, value)
	return self
}

func (self SpanElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self SpanElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *SpanElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *SpanElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *SpanElement) WithId(value string) *SpanElement {
	self.element.WithId(value)
	return self
}

func (self SpanElement) HasId() bool {
	return self.element.HasId()
}

func (self SpanElement) GetId() string {
	return self.element.GetId()
}

func (self *SpanElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *SpanElement) DelId() {
	self.element.DelId()
}

func (self *SpanElement) WithClass(classes ...string) *SpanElement {
	self.element.WithClass(classes...)
	return self
}

func (self SpanElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self SpanElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *SpanElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *SpanElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *SpanElement) WithStyles(styles ...maps.KeyValue[string, string]) *SpanElement {
	self.element.WithStyles(styles...)
	return self
}

func (self SpanElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *SpanElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *SpanElement) WithStyle(name string, value string) *SpanElement {
	self.element.WithStyle(name, value)
	return self
}

func (self SpanElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self SpanElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *SpanElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *SpanElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self SpanElement) Count() int {
	return self.element.Count()
}

func (self SpanElement) Children() []Node {
	return self.element.Children()
}

func (self *SpanElement) Push(children ...any) *SpanElement {
	self.element.Push(children...)
	return self
}

func (self *SpanElement) Pop() *SpanElement {
	self.element.Pop()
	return self
}

func (self SpanElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self SpanElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *SpanElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *SpanElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
