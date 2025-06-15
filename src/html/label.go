package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
type LabelElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/label
func Label(children ...any) *LabelElement {
	return &LabelElement{Elem("label").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/for
func (self *LabelElement) WithFor(value string) *LabelElement {
	return self.WithAttr("for", value)
}

func (self LabelElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *LabelElement) WithAttr(name string, value string) *LabelElement {
	self.element.WithAttr(name, value)
	return self
}

func (self LabelElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self LabelElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *LabelElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *LabelElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *LabelElement) WithId(value string) *LabelElement {
	self.element.WithId(value)
	return self
}

func (self LabelElement) HasId() bool {
	return self.element.HasId()
}

func (self LabelElement) GetId() string {
	return self.element.GetId()
}

func (self *LabelElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *LabelElement) DelId() {
	self.element.DelId()
}

func (self *LabelElement) WithClass(classes ...string) *LabelElement {
	self.element.WithClass(classes...)
	return self
}

func (self LabelElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self LabelElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *LabelElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *LabelElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *LabelElement) WithStyles(styles ...maps.KeyValue[string, string]) *LabelElement {
	self.element.WithStyles(styles...)
	return self
}

func (self LabelElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *LabelElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *LabelElement) WithStyle(name string, value string) *LabelElement {
	self.element.WithStyle(name, value)
	return self
}

func (self LabelElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self LabelElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *LabelElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *LabelElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self LabelElement) Count() int {
	return self.element.Count()
}

func (self LabelElement) Children() []Node {
	return self.element.Children()
}

func (self *LabelElement) Push(children ...any) *LabelElement {
	self.element.Push(children...)
	return self
}

func (self *LabelElement) Pop() *LabelElement {
	self.element.Pop()
	return self
}

func (self LabelElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self LabelElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *LabelElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *LabelElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
