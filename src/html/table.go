package html

import (
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/table
type TableElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/table
func Table(children ...any) *TableElement {
	return &TableElement{Elem("table").Push(children...)}
}

func (self TableElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TableElement) WithAttr(name string, value string) *TableElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TableElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TableElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TableElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TableElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TableElement) WithId(value string) *TableElement {
	self.element.WithId(value)
	return self
}

func (self TableElement) HasId() bool {
	return self.element.HasId()
}

func (self TableElement) GetId() string {
	return self.element.GetId()
}

func (self *TableElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TableElement) DelId() {
	self.element.DelId()
}

func (self *TableElement) WithClass(classes ...string) *TableElement {
	self.element.WithClass(classes...)
	return self
}

func (self TableElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TableElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TableElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TableElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TableElement) WithStyles(styles ...maps.KeyValue[string, string]) *TableElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TableElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TableElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TableElement) WithStyle(name string, value string) *TableElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TableElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TableElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TableElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TableElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TableElement) Count() int {
	return self.element.Count()
}

func (self TableElement) Children() []Node {
	return self.element.Children()
}

func (self *TableElement) Push(children ...any) *TableElement {
	self.element.Push(children...)
	return self
}

func (self *TableElement) Pop() *TableElement {
	self.element.Pop()
	return self
}

func (self TableElement) Render() []byte {
	return self.element.Render()
}

func (self TableElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *TableElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TableElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
