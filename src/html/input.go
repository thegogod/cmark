package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input
type InputElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#input_types
func Input(kind string) *InputElement {
	self := &InputElement{Elem("input").Void()}
	return self.WithType(kind)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#input_types
func (self *InputElement) WithType(value string) *InputElement {
	return self.WithAttr("type", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *InputElement) WithName(value string) *InputElement {
	return self.WithAttr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *InputElement) WithDisabled() *InputElement {
	return self.WithAttr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *InputElement) WithEnabled() *InputElement {
	return self.WithAttr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *InputElement) WithForm(formId string) *InputElement {
	return self.WithAttr("form", formId)
}

func (self InputElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *InputElement) WithAttr(name string, value string) *InputElement {
	self.element.WithAttr(name, value)
	return self
}

func (self InputElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self InputElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *InputElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *InputElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *InputElement) WithId(value string) *InputElement {
	self.element.WithId(value)
	return self
}

func (self InputElement) HasId() bool {
	return self.element.HasId()
}

func (self InputElement) GetId() string {
	return self.element.GetId()
}

func (self *InputElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *InputElement) DelId() {
	self.element.DelId()
}

func (self *InputElement) WithClass(classes ...string) *InputElement {
	self.element.WithClass(classes...)
	return self
}

func (self InputElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self InputElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *InputElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *InputElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *InputElement) WithStyles(styles ...maps.KeyValue[string, string]) *InputElement {
	self.element.WithStyles(styles...)
	return self
}

func (self InputElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *InputElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *InputElement) WithStyle(name string, value string) *InputElement {
	self.element.WithStyle(name, value)
	return self
}

func (self InputElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self InputElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *InputElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *InputElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self InputElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self InputElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *InputElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *InputElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
