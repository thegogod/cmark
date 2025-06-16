package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/div
type DivElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/div
func Div(children ...any) *DivElement {
	return &DivElement{Elem("div").Push(children...)}
}

func (self DivElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *DivElement) WithAttr(name string, value string) *DivElement {
	self.element.WithAttr(name, value)
	return self
}

func (self DivElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self DivElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *DivElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *DivElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *DivElement) WithId(value string) *DivElement {
	self.element.WithId(value)
	return self
}

func (self DivElement) HasId() bool {
	return self.element.HasId()
}

func (self DivElement) GetId() string {
	return self.element.GetId()
}

func (self *DivElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *DivElement) DelId() {
	self.element.DelId()
}

func (self *DivElement) WithClass(classes ...string) *DivElement {
	self.element.WithClass(classes...)
	return self
}

func (self DivElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self DivElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *DivElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *DivElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *DivElement) WithStyles(styles ...maps.KeyValue[string, string]) *DivElement {
	self.element.WithStyles(styles...)
	return self
}

func (self DivElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *DivElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *DivElement) WithStyle(name string, value string) *DivElement {
	self.element.WithStyle(name, value)
	return self
}

func (self DivElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self DivElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *DivElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *DivElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self DivElement) Count() int {
	return self.element.Count()
}

func (self DivElement) Children() []Node {
	return self.element.Children()
}

func (self *DivElement) Push(children ...any) *DivElement {
	self.element.Push(children...)
	return self
}

func (self *DivElement) Pop() *DivElement {
	self.element.Pop()
	return self
}

func (self DivElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self DivElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *DivElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *DivElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self DivElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self DivElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
