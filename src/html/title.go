package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/title
type TitleElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/title
func Title(children ...any) *TitleElement {
	return &TitleElement{Elem("title").Push(children...)}
}

func (self TitleElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TitleElement) WithAttr(name string, value string) *TitleElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TitleElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TitleElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TitleElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TitleElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TitleElement) WithId(value string) *TitleElement {
	self.element.WithId(value)
	return self
}

func (self TitleElement) HasId() bool {
	return self.element.HasId()
}

func (self TitleElement) GetId() string {
	return self.element.GetId()
}

func (self *TitleElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TitleElement) DelId() {
	self.element.DelId()
}

func (self *TitleElement) WithClass(classes ...string) *TitleElement {
	self.element.WithClass(classes...)
	return self
}

func (self TitleElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TitleElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TitleElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TitleElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TitleElement) WithStyles(styles ...maps.KeyValue[string, string]) *TitleElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TitleElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TitleElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TitleElement) WithStyle(name string, value string) *TitleElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TitleElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TitleElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TitleElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TitleElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TitleElement) Count() int {
	return self.element.Count()
}

func (self TitleElement) Children() []Node {
	return self.element.Children()
}

func (self *TitleElement) Push(children ...any) *TitleElement {
	self.element.Push(children...)
	return self
}

func (self *TitleElement) Pop() *TitleElement {
	self.element.Pop()
	return self
}

func (self TitleElement) Render() []byte {
	return self.element.Render()
}

func (self TitleElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *TitleElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TitleElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
