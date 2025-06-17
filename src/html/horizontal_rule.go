package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hr
type HorizontalRuleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/hr
func Hr() *HorizontalRuleElement {
	return &HorizontalRuleElement{Elem("hr").Void()}
}

func (self HorizontalRuleElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *HorizontalRuleElement) WithAttr(name string, value string) *HorizontalRuleElement {
	self.element.WithAttr(name, value)
	return self
}

func (self HorizontalRuleElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self HorizontalRuleElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *HorizontalRuleElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *HorizontalRuleElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *HorizontalRuleElement) WithId(value string) *HorizontalRuleElement {
	self.element.WithId(value)
	return self
}

func (self HorizontalRuleElement) HasId() bool {
	return self.element.HasId()
}

func (self HorizontalRuleElement) GetId() string {
	return self.element.GetId()
}

func (self *HorizontalRuleElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *HorizontalRuleElement) DelId() {
	self.element.DelId()
}

func (self *HorizontalRuleElement) WithClass(classes ...string) *HorizontalRuleElement {
	self.element.WithClass(classes...)
	return self
}

func (self HorizontalRuleElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self HorizontalRuleElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *HorizontalRuleElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *HorizontalRuleElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *HorizontalRuleElement) WithStyles(styles ...maps.KeyValue[string, string]) *HorizontalRuleElement {
	self.element.WithStyles(styles...)
	return self
}

func (self HorizontalRuleElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *HorizontalRuleElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *HorizontalRuleElement) WithStyle(name string, value string) *HorizontalRuleElement {
	self.element.WithStyle(name, value)
	return self
}

func (self HorizontalRuleElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self HorizontalRuleElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *HorizontalRuleElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *HorizontalRuleElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self HorizontalRuleElement) Render() []byte {
	return self.element.Render()
}

func (self HorizontalRuleElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *HorizontalRuleElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *HorizontalRuleElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
