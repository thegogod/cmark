package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/i
type ItalicElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/i
func I(children ...any) *ItalicElement {
	return &ItalicElement{Elem("i").Push(children...)}
}

func (self ItalicElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *ItalicElement) WithAttr(name string, value string) *ItalicElement {
	self.element.WithAttr(name, value)
	return self
}

func (self ItalicElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self ItalicElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *ItalicElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *ItalicElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *ItalicElement) WithId(value string) *ItalicElement {
	self.element.WithId(value)
	return self
}

func (self ItalicElement) HasId() bool {
	return self.element.HasId()
}

func (self ItalicElement) GetId() string {
	return self.element.GetId()
}

func (self *ItalicElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *ItalicElement) DelId() {
	self.element.DelId()
}

func (self *ItalicElement) WithClass(classes ...string) *ItalicElement {
	self.element.WithClass(classes...)
	return self
}

func (self ItalicElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self ItalicElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *ItalicElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *ItalicElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *ItalicElement) WithStyles(styles ...maps.KeyValue[string, string]) *ItalicElement {
	self.element.WithStyles(styles...)
	return self
}

func (self ItalicElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ItalicElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *ItalicElement) WithStyle(name string, value string) *ItalicElement {
	self.element.WithStyle(name, value)
	return self
}

func (self ItalicElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self ItalicElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *ItalicElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *ItalicElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self ItalicElement) Count() int {
	return self.element.Count()
}

func (self ItalicElement) Children() []Node {
	return self.element.Children()
}

func (self *ItalicElement) Push(children ...any) *ItalicElement {
	self.element.Push(children...)
	return self
}

func (self *ItalicElement) Pop() *ItalicElement {
	self.element.Pop()
	return self
}

func (self ItalicElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self ItalicElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *ItalicElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *ItalicElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self ItalicElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self ItalicElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
