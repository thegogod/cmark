package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template
type TemplateElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template
func Template(children ...any) *TemplateElement {
	return &TemplateElement{Elem("template").Push(children...)}
}

func (self TemplateElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TemplateElement) WithAttr(name string, value string) *TemplateElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TemplateElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TemplateElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TemplateElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TemplateElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TemplateElement) WithId(value string) *TemplateElement {
	self.element.WithId(value)
	return self
}

func (self TemplateElement) HasId() bool {
	return self.element.HasId()
}

func (self TemplateElement) GetId() string {
	return self.element.GetId()
}

func (self *TemplateElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TemplateElement) DelId() {
	self.element.DelId()
}

func (self *TemplateElement) WithClass(classes ...string) *TemplateElement {
	self.element.WithClass(classes...)
	return self
}

func (self TemplateElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TemplateElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TemplateElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TemplateElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TemplateElement) WithStyles(styles ...maps.KeyValue[string, string]) *TemplateElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TemplateElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TemplateElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TemplateElement) WithStyle(name string, value string) *TemplateElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TemplateElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TemplateElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TemplateElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TemplateElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TemplateElement) Count() int {
	return self.element.Count()
}

func (self TemplateElement) Children() []Node {
	return self.element.Children()
}

func (self *TemplateElement) Push(children ...any) *TemplateElement {
	self.element.Push(children...)
	return self
}

func (self *TemplateElement) Pop() *TemplateElement {
	self.element.Pop()
	return self
}

func (self TemplateElement) Render() []byte {
	return self.element.Render()
}

func (self TemplateElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *TemplateElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TemplateElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
