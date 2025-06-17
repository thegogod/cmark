package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/mark
type MarkElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/mark
func Mark(children ...any) *MarkElement {
	return &MarkElement{Elem("mark").Push(children...)}
}

func (self MarkElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *MarkElement) WithAttr(name string, value string) *MarkElement {
	self.element.WithAttr(name, value)
	return self
}

func (self MarkElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self MarkElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *MarkElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *MarkElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *MarkElement) WithId(value string) *MarkElement {
	self.element.WithId(value)
	return self
}

func (self MarkElement) HasId() bool {
	return self.element.HasId()
}

func (self MarkElement) GetId() string {
	return self.element.GetId()
}

func (self *MarkElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *MarkElement) DelId() {
	self.element.DelId()
}

func (self *MarkElement) WithClass(classes ...string) *MarkElement {
	self.element.WithClass(classes...)
	return self
}

func (self MarkElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self MarkElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *MarkElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *MarkElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *MarkElement) WithStyles(styles ...maps.KeyValue[string, string]) *MarkElement {
	self.element.WithStyles(styles...)
	return self
}

func (self MarkElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *MarkElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *MarkElement) WithStyle(name string, value string) *MarkElement {
	self.element.WithStyle(name, value)
	return self
}

func (self MarkElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self MarkElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *MarkElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *MarkElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self MarkElement) Count() int {
	return self.element.Count()
}

func (self MarkElement) Children() []Node {
	return self.element.Children()
}

func (self *MarkElement) Push(children ...any) *MarkElement {
	self.element.Push(children...)
	return self
}

func (self *MarkElement) Pop() *MarkElement {
	self.element.Pop()
	return self
}

func (self MarkElement) Render() []byte {
	return self.element.Render()
}

func (self MarkElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *MarkElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *MarkElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
