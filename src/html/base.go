package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base
type BaseElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base
func Base() *BaseElement {
	return &BaseElement{Elem("base").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base#href
func (self *BaseElement) WithHref(value string) *BaseElement {
	return self.WithAttr("href", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/base#target
func (self *BaseElement) WithTarget(value string) *BaseElement {
	return self.WithAttr("target", value)
}

func (self BaseElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *BaseElement) WithAttr(name string, value string) *BaseElement {
	self.element.WithAttr(name, value)
	return self
}

func (self BaseElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self BaseElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *BaseElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *BaseElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *BaseElement) WithId(value string) *BaseElement {
	self.element.WithId(value)
	return self
}

func (self BaseElement) HasId() bool {
	return self.element.HasId()
}

func (self BaseElement) GetId() string {
	return self.element.GetId()
}

func (self *BaseElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *BaseElement) DelId() {
	self.element.DelId()
}

func (self *BaseElement) WithClass(classes ...string) *BaseElement {
	self.element.WithClass(classes...)
	return self
}

func (self BaseElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self BaseElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *BaseElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *BaseElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *BaseElement) WithStyles(styles ...maps.KeyValue[string, string]) *BaseElement {
	self.element.WithStyles(styles...)
	return self
}

func (self BaseElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *BaseElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *BaseElement) WithStyle(name string, value string) *BaseElement {
	self.element.WithStyle(name, value)
	return self
}

func (self BaseElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self BaseElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *BaseElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *BaseElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self BaseElement) Render() []byte {
	return self.element.Render()
}

func (self BaseElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *BaseElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *BaseElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
