package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
type StyleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/style
func Style(children ...any) *StyleElement {
	return &StyleElement{Elem("style").WithAttr("type", "text/css").Push(children...)}
}

func (self StyleElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *StyleElement) WithAttr(name string, value string) *StyleElement {
	self.element.WithAttr(name, value)
	return self
}

func (self StyleElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self StyleElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *StyleElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *StyleElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *StyleElement) WithId(value string) *StyleElement {
	self.element.WithId(value)
	return self
}

func (self StyleElement) HasId() bool {
	return self.element.HasId()
}

func (self StyleElement) GetId() string {
	return self.element.GetId()
}

func (self *StyleElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *StyleElement) DelId() {
	self.element.DelId()
}

func (self *StyleElement) WithClass(classes ...string) *StyleElement {
	self.element.WithClass(classes...)
	return self
}

func (self StyleElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self StyleElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *StyleElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *StyleElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *StyleElement) WithStyles(styles ...maps.KeyValue[string, string]) *StyleElement {
	self.element.WithStyles(styles...)
	return self
}

func (self StyleElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *StyleElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *StyleElement) WithStyle(name string, value string) *StyleElement {
	self.element.WithStyle(name, value)
	return self
}

func (self StyleElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self StyleElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *StyleElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *StyleElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self StyleElement) Count() int {
	return self.element.Count()
}

func (self StyleElement) Children() []Node {
	return self.element.Children()
}

func (self *StyleElement) Push(children ...any) *StyleElement {
	self.element.Push(children...)
	return self
}

func (self *StyleElement) Pop() *StyleElement {
	self.element.Pop()
	return self
}

func (self StyleElement) Render() []byte {
	return self.element.Render()
}

func (self StyleElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *StyleElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *StyleElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
