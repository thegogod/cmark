package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
type HeadElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/head
func Head(children ...any) *HeadElement {
	return &HeadElement{Elem("head").Push(children...)}
}

func (self HeadElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *HeadElement) WithAttr(name string, value string) *HeadElement {
	self.element.WithAttr(name, value)
	return self
}

func (self HeadElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self HeadElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *HeadElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *HeadElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *HeadElement) WithId(value string) *HeadElement {
	self.element.WithId(value)
	return self
}

func (self HeadElement) HasId() bool {
	return self.element.HasId()
}

func (self HeadElement) GetId() string {
	return self.element.GetId()
}

func (self *HeadElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *HeadElement) DelId() {
	self.element.DelId()
}

func (self *HeadElement) WithClass(classes ...string) *HeadElement {
	self.element.WithClass(classes...)
	return self
}

func (self HeadElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self HeadElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *HeadElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *HeadElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *HeadElement) WithStyles(styles ...maps.KeyValue[string, string]) *HeadElement {
	self.element.WithStyles(styles...)
	return self
}

func (self HeadElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *HeadElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *HeadElement) WithStyle(name string, value string) *HeadElement {
	self.element.WithStyle(name, value)
	return self
}

func (self HeadElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self HeadElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *HeadElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *HeadElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self HeadElement) Count() int {
	return self.element.Count()
}

func (self HeadElement) Children() []Node {
	return self.element.Children()
}

func (self *HeadElement) Push(children ...any) *HeadElement {
	self.element.Push(children...)
	return self
}

func (self *HeadElement) Pop() *HeadElement {
	self.element.Pop()
	return self
}

func (self HeadElement) Render() []byte {
	return self.element.Render()
}

func (self HeadElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *HeadElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *HeadElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
