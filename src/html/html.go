package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html
type HtmlElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html
func Html(children ...any) *HtmlElement {
	return &HtmlElement{Elem("html").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/html#accessibility
func (self *HtmlElement) WithLang(value string) *HtmlElement {
	return self.WithAttr("lang", value)
}

func (self HtmlElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *HtmlElement) WithAttr(name string, value string) *HtmlElement {
	self.element.WithAttr(name, value)
	return self
}

func (self HtmlElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self HtmlElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *HtmlElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *HtmlElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *HtmlElement) WithId(value string) *HtmlElement {
	self.element.WithId(value)
	return self
}

func (self HtmlElement) HasId() bool {
	return self.element.HasId()
}

func (self HtmlElement) GetId() string {
	return self.element.GetId()
}

func (self *HtmlElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *HtmlElement) DelId() {
	self.element.DelId()
}

func (self *HtmlElement) WithClass(classes ...string) *HtmlElement {
	self.element.WithClass(classes...)
	return self
}

func (self HtmlElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self HtmlElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *HtmlElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *HtmlElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *HtmlElement) WithStyles(styles ...maps.KeyValue[string, string]) *HtmlElement {
	self.element.WithStyles(styles...)
	return self
}

func (self HtmlElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *HtmlElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *HtmlElement) WithStyle(name string, value string) *HtmlElement {
	self.element.WithStyle(name, value)
	return self
}

func (self HtmlElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self HtmlElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *HtmlElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *HtmlElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self HtmlElement) Count() int {
	return self.element.Count()
}

func (self HtmlElement) Children() []Node {
	return self.element.Children()
}

func (self *HtmlElement) Push(children ...any) *HtmlElement {
	self.element.Push(children...)
	return self
}

func (self *HtmlElement) Pop() *HtmlElement {
	self.element.Pop()
	return self
}

func (self HtmlElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self HtmlElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *HtmlElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *HtmlElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
