package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox
type CheckBoxInputElement struct {
	element *InputElement
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox
func CheckBoxInput() *CheckBoxInputElement {
	return &CheckBoxInputElement{Input("checkbox")}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *CheckBoxInputElement) WithName(value string) *CheckBoxInputElement {
	return self.WithAttr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *CheckBoxInputElement) WithDisabled() *CheckBoxInputElement {
	return self.WithAttr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *CheckBoxInputElement) WithEnabled() *CheckBoxInputElement {
	return self.WithAttr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *CheckBoxInputElement) WithForm(formId string) *CheckBoxInputElement {
	return self.WithAttr("form", formId)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox#value
func (self *CheckBoxInputElement) WithValue(value string) *CheckBoxInputElement {
	return self.WithAttr("value", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/checkbox#checked
func (self *CheckBoxInputElement) WithChecked(value bool) *CheckBoxInputElement {
	if value {
		return self.WithAttr("checked", "true")
	}

	self.DelAttr("checked")
	return self
}

func (self CheckBoxInputElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *CheckBoxInputElement) WithAttr(name string, value string) *CheckBoxInputElement {
	self.element.WithAttr(name, value)
	return self
}

func (self CheckBoxInputElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self CheckBoxInputElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *CheckBoxInputElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *CheckBoxInputElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *CheckBoxInputElement) WithId(value string) *CheckBoxInputElement {
	self.element.WithId(value)
	return self
}

func (self CheckBoxInputElement) HasId() bool {
	return self.element.HasId()
}

func (self CheckBoxInputElement) GetId() string {
	return self.element.GetId()
}

func (self *CheckBoxInputElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *CheckBoxInputElement) DelId() {
	self.element.DelId()
}

func (self *CheckBoxInputElement) WithClass(classes ...string) *CheckBoxInputElement {
	self.element.WithClass(classes...)
	return self
}

func (self CheckBoxInputElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self CheckBoxInputElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *CheckBoxInputElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *CheckBoxInputElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *CheckBoxInputElement) WithStyles(styles ...maps.KeyValue[string, string]) *CheckBoxInputElement {
	self.element.WithStyles(styles...)
	return self
}

func (self CheckBoxInputElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *CheckBoxInputElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *CheckBoxInputElement) WithStyle(name string, value string) *CheckBoxInputElement {
	self.element.WithStyle(name, value)
	return self
}

func (self CheckBoxInputElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self CheckBoxInputElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *CheckBoxInputElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *CheckBoxInputElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self CheckBoxInputElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self CheckBoxInputElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *CheckBoxInputElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *CheckBoxInputElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}

func (self CheckBoxInputElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self CheckBoxInputElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
