package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/strong
type StrongElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/strong
func Strong(children ...any) *StrongElement {
	return &StrongElement{Elem("strong").Push(children...)}
}

func (self StrongElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *StrongElement) WithAttr(name string, value string) *StrongElement {
	self.element.WithAttr(name, value)
	return self
}

func (self StrongElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self StrongElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *StrongElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *StrongElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *StrongElement) WithId(value string) *StrongElement {
	self.element.WithId(value)
	return self
}

func (self StrongElement) HasId() bool {
	return self.element.HasId()
}

func (self StrongElement) GetId() string {
	return self.element.GetId()
}

func (self *StrongElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *StrongElement) DelId() {
	self.element.DelId()
}

func (self *StrongElement) WithClass(classes ...string) *StrongElement {
	self.element.WithClass(classes...)
	return self
}

func (self StrongElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self StrongElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *StrongElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *StrongElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *StrongElement) WithStyles(styles ...maps.KeyValue[string, string]) *StrongElement {
	self.element.WithStyles(styles...)
	return self
}

func (self StrongElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *StrongElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *StrongElement) WithStyle(name string, value string) *StrongElement {
	self.element.WithStyle(name, value)
	return self
}

func (self StrongElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self StrongElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *StrongElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *StrongElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self StrongElement) Count() int {
	return self.element.Count()
}

func (self StrongElement) Children() []Node {
	return self.element.Children()
}

func (self *StrongElement) Push(children ...any) *StrongElement {
	self.element.Push(children...)
	return self
}

func (self *StrongElement) Pop() *StrongElement {
	self.element.Pop()
	return self
}

func (self StrongElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self StrongElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *StrongElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *StrongElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self StrongElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self StrongElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
