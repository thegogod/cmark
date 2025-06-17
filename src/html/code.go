package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/code
type CodeElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/code
func Code(children ...any) *CodeElement {
	return &CodeElement{Elem("code").Push(children...)}
}

func (self CodeElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *CodeElement) WithAttr(name string, value string) *CodeElement {
	self.element.WithAttr(name, value)
	return self
}

func (self CodeElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self CodeElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *CodeElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *CodeElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *CodeElement) WithId(value string) *CodeElement {
	self.element.WithId(value)
	return self
}

func (self CodeElement) HasId() bool {
	return self.element.HasId()
}

func (self CodeElement) GetId() string {
	return self.element.GetId()
}

func (self *CodeElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *CodeElement) DelId() {
	self.element.DelId()
}

func (self *CodeElement) WithClass(classes ...string) *CodeElement {
	self.element.WithClass(classes...)
	return self
}

func (self CodeElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self CodeElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *CodeElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *CodeElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *CodeElement) WithStyles(styles ...maps.KeyValue[string, string]) *CodeElement {
	self.element.WithStyles(styles...)
	return self
}

func (self CodeElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *CodeElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *CodeElement) WithStyle(name string, value string) *CodeElement {
	self.element.WithStyle(name, value)
	return self
}

func (self CodeElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self CodeElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *CodeElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *CodeElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self CodeElement) Count() int {
	return self.element.Count()
}

func (self CodeElement) Children() []Node {
	return self.element.Children()
}

func (self *CodeElement) Push(children ...any) *CodeElement {
	self.element.Push(children...)
	return self
}

func (self *CodeElement) Pop() *CodeElement {
	self.element.Pop()
	return self
}

func (self CodeElement) Render() []byte {
	return self.element.Render()
}

func (self CodeElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *CodeElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *CodeElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
