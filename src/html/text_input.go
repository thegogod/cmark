package html

import (
	"strconv"

	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text
type TextInputElement struct {
	element *InputElement
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text
func TextInput() *TextInputElement {
	return &TextInputElement{Input("text")}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *TextInputElement) WithName(value string) *TextInputElement {
	return self.WithAttr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *TextInputElement) WithDisabled() *TextInputElement {
	return self.WithAttr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *TextInputElement) WithEnabled() *TextInputElement {
	return self.WithAttr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *TextInputElement) WithForm(formId string) *TextInputElement {
	return self.WithAttr("form", formId)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#value
func (self *TextInputElement) WithValue(value string) *TextInputElement {
	return self.WithAttr("value", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#maxlength
func (self *TextInputElement) WithMaxLength(value int) *TextInputElement {
	return self.WithAttr("maxlength", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#minlength
func (self *TextInputElement) WithMinLength(value int) *TextInputElement {
	return self.WithAttr("minlength", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#pattern
func (self *TextInputElement) WithPattern(value string) *TextInputElement {
	return self.WithAttr("pattern", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#placeholder
func (self *TextInputElement) WithPlaceholder(value string) *TextInputElement {
	return self.WithAttr("placeholder", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#readonly
func (self *TextInputElement) WithReadonly() *TextInputElement {
	return self.WithAttr("readonly", "")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#size
func (self *TextInputElement) WithSize(value int) *TextInputElement {
	return self.WithAttr("size", strconv.Itoa(value))
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/text#spellcheck
func (self *TextInputElement) WithSpellcheck(value bool) *TextInputElement {
	return self.WithAttr("spellcheck", strconv.FormatBool(value))
}

func (self TextInputElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *TextInputElement) WithAttr(name string, value string) *TextInputElement {
	self.element.WithAttr(name, value)
	return self
}

func (self TextInputElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self TextInputElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *TextInputElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *TextInputElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *TextInputElement) WithId(value string) *TextInputElement {
	self.element.WithId(value)
	return self
}

func (self TextInputElement) HasId() bool {
	return self.element.HasId()
}

func (self TextInputElement) GetId() string {
	return self.element.GetId()
}

func (self *TextInputElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *TextInputElement) DelId() {
	self.element.DelId()
}

func (self *TextInputElement) WithClass(classes ...string) *TextInputElement {
	self.element.WithClass(classes...)
	return self
}

func (self TextInputElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self TextInputElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *TextInputElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *TextInputElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *TextInputElement) WithStyles(styles ...maps.KeyValue[string, string]) *TextInputElement {
	self.element.WithStyles(styles...)
	return self
}

func (self TextInputElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *TextInputElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *TextInputElement) WithStyle(name string, value string) *TextInputElement {
	self.element.WithStyle(name, value)
	return self
}

func (self TextInputElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self TextInputElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *TextInputElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *TextInputElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self TextInputElement) Render() []byte {
	return self.element.Render()
}

func (self TextInputElement) RenderPretty(indent string) []byte {
	return self.element.RenderPretty(indent)
}

func (self *TextInputElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *TextInputElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
