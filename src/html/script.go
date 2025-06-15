package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
type ScriptElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script
func Script(children ...any) *ScriptElement {
	return &ScriptElement{Elem("script").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script#src
func (self *ScriptElement) WithSrc(value string) *ScriptElement {
	return self.WithAttr("src", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script/type
func (self *ScriptElement) WithType(value string) *ScriptElement {
	return self.WithAttr("type", value)
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script#async
func (self *ScriptElement) WithAsync() *ScriptElement {
	return self.WithAttr("async", "")
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/script#defer
func (self *ScriptElement) WithDefer() *ScriptElement {
	return self.WithAttr("defer", "")
}

func (self ScriptElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *ScriptElement) WithAttr(name string, value string) *ScriptElement {
	self.element.WithAttr(name, value)
	return self
}

func (self ScriptElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self ScriptElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *ScriptElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *ScriptElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *ScriptElement) WithId(value string) *ScriptElement {
	self.element.WithId(value)
	return self
}

func (self ScriptElement) HasId() bool {
	return self.element.HasId()
}

func (self ScriptElement) GetId() string {
	return self.element.GetId()
}

func (self *ScriptElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *ScriptElement) DelId() {
	self.element.DelId()
}

func (self *ScriptElement) WithClass(classes ...string) *ScriptElement {
	self.element.WithClass(classes...)
	return self
}

func (self ScriptElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self ScriptElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *ScriptElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *ScriptElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *ScriptElement) WithStyles(styles ...maps.KeyValue[string, string]) *ScriptElement {
	self.element.WithStyles(styles...)
	return self
}

func (self ScriptElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ScriptElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *ScriptElement) WithStyle(name string, value string) *ScriptElement {
	self.element.WithStyle(name, value)
	return self
}

func (self ScriptElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self ScriptElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *ScriptElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *ScriptElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self ScriptElement) Count() int {
	return self.element.Count()
}

func (self ScriptElement) Children() []Node {
	return self.element.Children()
}

func (self *ScriptElement) Push(children ...any) *ScriptElement {
	self.element.Push(children...)
	return self
}

func (self *ScriptElement) Pop() *ScriptElement {
	self.element.Pop()
	return self
}

func (self ScriptElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self ScriptElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *ScriptElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *ScriptElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
