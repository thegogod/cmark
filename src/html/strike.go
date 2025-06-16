package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/s
type StrikeElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/s
func S(children ...any) *StrikeElement {
	return &StrikeElement{Elem("s").Push(children...)}
}

func (self StrikeElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *StrikeElement) WithAttr(name string, value string) *StrikeElement {
	self.element.WithAttr(name, value)
	return self
}

func (self StrikeElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self StrikeElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *StrikeElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *StrikeElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *StrikeElement) WithId(value string) *StrikeElement {
	self.element.WithId(value)
	return self
}

func (self StrikeElement) HasId() bool {
	return self.element.HasId()
}

func (self StrikeElement) GetId() string {
	return self.element.GetId()
}

func (self *StrikeElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *StrikeElement) DelId() {
	self.element.DelId()
}

func (self *StrikeElement) WithClass(classes ...string) *StrikeElement {
	self.element.WithClass(classes...)
	return self
}

func (self StrikeElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self StrikeElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *StrikeElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *StrikeElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *StrikeElement) WithStyles(styles ...maps.KeyValue[string, string]) *StrikeElement {
	self.element.WithStyles(styles...)
	return self
}

func (self StrikeElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *StrikeElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *StrikeElement) WithStyle(name string, value string) *StrikeElement {
	self.element.WithStyle(name, value)
	return self
}

func (self StrikeElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self StrikeElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *StrikeElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *StrikeElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self StrikeElement) Count() int {
	return self.element.Count()
}

func (self StrikeElement) Children() []Node {
	return self.element.Children()
}

func (self *StrikeElement) Push(children ...any) *StrikeElement {
	self.element.Push(children...)
	return self
}

func (self *StrikeElement) Pop() *StrikeElement {
	self.element.Pop()
	return self
}

func (self StrikeElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self StrikeElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *StrikeElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *StrikeElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self StrikeElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self StrikeElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
