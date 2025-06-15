package html

import (
	"strconv"

	"github.com/thegogod/cmark/ast"

	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file
type FileInputElement struct {
	element *InputElement
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file
func FileInput() *FileInputElement {
	return &FileInputElement{Input("file")}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#name
func (self *FileInputElement) WithName(value string) *FileInputElement {
	return self.WithAttr("name", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *FileInputElement) WithDisabled() *FileInputElement {
	return self.WithAttr("disabled", "true")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#disabled
func (self *FileInputElement) WithEnabled() *FileInputElement {
	return self.WithAttr("disabled", "false")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#form
func (self *FileInputElement) WithForm(formId string) *FileInputElement {
	return self.WithAttr("form", formId)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#value
func (self *FileInputElement) WithValue(value string) *FileInputElement {
	return self.WithAttr("value", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#accept
func (self *FileInputElement) WithAccept(value string) *FileInputElement {
	return self.WithAttr("accept", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#capture
func (self *FileInputElement) WithCapture(value string) *FileInputElement {
	return self.WithAttr("capture", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input/file#multiple
func (self *FileInputElement) WithMultiple(value bool) *FileInputElement {
	return self.WithAttr("multiple", strconv.FormatBool(value))
}

func (self FileInputElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *FileInputElement) WithAttr(name string, value string) *FileInputElement {
	self.element.WithAttr(name, value)
	return self
}

func (self FileInputElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self FileInputElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *FileInputElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *FileInputElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *FileInputElement) WithId(value string) *FileInputElement {
	self.element.WithId(value)
	return self
}

func (self FileInputElement) HasId() bool {
	return self.element.HasId()
}

func (self FileInputElement) GetId() string {
	return self.element.GetId()
}

func (self *FileInputElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *FileInputElement) DelId() {
	self.element.DelId()
}

func (self *FileInputElement) WithClass(classes ...string) *FileInputElement {
	self.element.WithClass(classes...)
	return self
}

func (self FileInputElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self FileInputElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *FileInputElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *FileInputElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *FileInputElement) WithStyles(styles ...maps.KeyValue[string, string]) *FileInputElement {
	self.element.WithStyles(styles...)
	return self
}

func (self FileInputElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *FileInputElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *FileInputElement) WithStyle(name string, value string) *FileInputElement {
	self.element.WithStyle(name, value)
	return self
}

func (self FileInputElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self FileInputElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *FileInputElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *FileInputElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self FileInputElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self FileInputElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *FileInputElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *FileInputElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
