package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
type BodyElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/body
func Body(children ...any) *BodyElement {
	return &BodyElement{Elem("body").Push(children...)}
}

func (self BodyElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *BodyElement) WithAttr(name string, value string) *BodyElement {
	self.element.WithAttr(name, value)
	return self
}

func (self BodyElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self BodyElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *BodyElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *BodyElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *BodyElement) WithId(value string) *BodyElement {
	self.element.WithId(value)
	return self
}

func (self BodyElement) HasId() bool {
	return self.element.HasId()
}

func (self BodyElement) GetId() string {
	return self.element.GetId()
}

func (self *BodyElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *BodyElement) DelId() {
	self.element.DelId()
}

func (self *BodyElement) WithClass(classes ...string) *BodyElement {
	self.element.WithClass(classes...)
	return self
}

func (self BodyElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self BodyElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *BodyElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *BodyElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *BodyElement) WithStyles(styles ...maps.KeyValue[string, string]) *BodyElement {
	self.element.WithStyles(styles...)
	return self
}

func (self BodyElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *BodyElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *BodyElement) WithStyle(name string, value string) *BodyElement {
	self.element.WithStyle(name, value)
	return self
}

func (self BodyElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self BodyElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *BodyElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *BodyElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self BodyElement) Count() int {
	return self.element.Count()
}

func (self BodyElement) Children() []Node {
	return self.element.Children()
}

func (self *BodyElement) Push(children ...any) *BodyElement {
	self.element.Push(children...)
	return self
}

func (self *BodyElement) Pop() *BodyElement {
	self.element.Pop()
	return self
}

func (self BodyElement) Render() []byte {
	return self.element.Render()
}

func (self BodyElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *BodyElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *BodyElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
