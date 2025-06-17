package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/br
type BreakLineElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/br
func Br() *BreakLineElement {
	return &BreakLineElement{Elem("br").Void()}
}

func (self BreakLineElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *BreakLineElement) WithAttr(name string, value string) *BreakLineElement {
	self.element.WithAttr(name, value)
	return self
}

func (self BreakLineElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self BreakLineElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *BreakLineElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *BreakLineElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *BreakLineElement) WithId(value string) *BreakLineElement {
	self.element.WithId(value)
	return self
}

func (self BreakLineElement) HasId() bool {
	return self.element.HasId()
}

func (self BreakLineElement) GetId() string {
	return self.element.GetId()
}

func (self *BreakLineElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *BreakLineElement) DelId() {
	self.element.DelId()
}

func (self *BreakLineElement) WithClass(classes ...string) *BreakLineElement {
	self.element.WithClass(classes...)
	return self
}

func (self BreakLineElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self BreakLineElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *BreakLineElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *BreakLineElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *BreakLineElement) WithStyles(styles ...maps.KeyValue[string, string]) *BreakLineElement {
	self.element.WithStyles(styles...)
	return self
}

func (self BreakLineElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *BreakLineElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *BreakLineElement) WithStyle(name string, value string) *BreakLineElement {
	self.element.WithStyle(name, value)
	return self
}

func (self BreakLineElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self BreakLineElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *BreakLineElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *BreakLineElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self BreakLineElement) Render() []byte {
	return self.element.Render()
}

func (self BreakLineElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *BreakLineElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *BreakLineElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
