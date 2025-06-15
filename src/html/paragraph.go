package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
type ParagraphElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/p
func P(children ...any) *ParagraphElement {
	return &ParagraphElement{Elem("p").Push(children...)}
}

func (self ParagraphElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *ParagraphElement) WithAttr(name string, value string) *ParagraphElement {
	self.element.WithAttr(name, value)
	return self
}

func (self ParagraphElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self ParagraphElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *ParagraphElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *ParagraphElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *ParagraphElement) WithId(value string) *ParagraphElement {
	self.element.WithId(value)
	return self
}

func (self ParagraphElement) HasId() bool {
	return self.element.HasId()
}

func (self ParagraphElement) GetId() string {
	return self.element.GetId()
}

func (self *ParagraphElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *ParagraphElement) DelId() {
	self.element.DelId()
}

func (self *ParagraphElement) WithClass(classes ...string) *ParagraphElement {
	self.element.WithClass(classes...)
	return self
}

func (self ParagraphElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self ParagraphElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *ParagraphElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *ParagraphElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *ParagraphElement) WithStyles(styles ...maps.KeyValue[string, string]) *ParagraphElement {
	self.element.WithStyles(styles...)
	return self
}

func (self ParagraphElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ParagraphElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *ParagraphElement) WithStyle(name string, value string) *ParagraphElement {
	self.element.WithStyle(name, value)
	return self
}

func (self ParagraphElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self ParagraphElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *ParagraphElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *ParagraphElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self ParagraphElement) Count() int {
	return self.element.Count()
}

func (self ParagraphElement) Children() []Node {
	return self.element.Children()
}

func (self *ParagraphElement) Push(children ...any) *ParagraphElement {
	self.element.Push(children...)
	return self
}

func (self *ParagraphElement) Pop() *ParagraphElement {
	self.element.Pop()
	return self
}

func (self ParagraphElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self ParagraphElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *ParagraphElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *ParagraphElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
