package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
type LinkElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link
func Link() *LinkElement {
	return &LinkElement{Elem("link").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/rel
func (self *LinkElement) WithRel(value string) *LinkElement {
	return self.WithAttr("rel", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/link#href
func (self *LinkElement) WithHref(value string) *LinkElement {
	return self.WithAttr("href", value)
}

func (self LinkElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *LinkElement) WithAttr(name string, value string) *LinkElement {
	self.element.WithAttr(name, value)
	return self
}

func (self LinkElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self LinkElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *LinkElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *LinkElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *LinkElement) WithId(value string) *LinkElement {
	self.element.WithId(value)
	return self
}

func (self LinkElement) HasId() bool {
	return self.element.HasId()
}

func (self LinkElement) GetId() string {
	return self.element.GetId()
}

func (self *LinkElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *LinkElement) DelId() {
	self.element.DelId()
}

func (self *LinkElement) WithClass(classes ...string) *LinkElement {
	self.element.WithClass(classes...)
	return self
}

func (self LinkElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self LinkElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *LinkElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *LinkElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *LinkElement) WithStyles(styles ...maps.KeyValue[string, string]) *LinkElement {
	self.element.WithStyles(styles...)
	return self
}

func (self LinkElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *LinkElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *LinkElement) WithStyle(name string, value string) *LinkElement {
	self.element.WithStyle(name, value)
	return self
}

func (self LinkElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self LinkElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *LinkElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *LinkElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self LinkElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self LinkElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *LinkElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *LinkElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self LinkElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self LinkElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
