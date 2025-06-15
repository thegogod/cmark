package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote
type BlockQuoteElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote
func BlockQuote(children ...any) *BlockQuoteElement {
	return &BlockQuoteElement{Elem("blockquote").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote#cite
func (self *BlockQuoteElement) WithCite(value string) *BlockQuoteElement {
	return self.WithAttr("cite", value)
}

func (self BlockQuoteElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *BlockQuoteElement) WithAttr(name string, value string) *BlockQuoteElement {
	self.element.WithAttr(name, value)
	return self
}

func (self BlockQuoteElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self BlockQuoteElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *BlockQuoteElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *BlockQuoteElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *BlockQuoteElement) WithId(value string) *BlockQuoteElement {
	self.element.WithId(value)
	return self
}

func (self BlockQuoteElement) HasId() bool {
	return self.element.HasId()
}

func (self BlockQuoteElement) GetId() string {
	return self.element.GetId()
}

func (self *BlockQuoteElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *BlockQuoteElement) DelId() {
	self.element.DelId()
}

func (self *BlockQuoteElement) WithClass(classes ...string) *BlockQuoteElement {
	self.element.WithClass(classes...)
	return self
}

func (self BlockQuoteElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self BlockQuoteElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *BlockQuoteElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *BlockQuoteElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *BlockQuoteElement) WithStyles(styles ...maps.KeyValue[string, string]) *BlockQuoteElement {
	self.element.WithStyles(styles...)
	return self
}

func (self BlockQuoteElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *BlockQuoteElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *BlockQuoteElement) WithStyle(name string, value string) *BlockQuoteElement {
	self.element.WithStyle(name, value)
	return self
}

func (self BlockQuoteElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self BlockQuoteElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *BlockQuoteElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *BlockQuoteElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self BlockQuoteElement) Count() int {
	return self.element.Count()
}

func (self BlockQuoteElement) Children() []Node {
	return self.element.Children()
}

func (self *BlockQuoteElement) Push(children ...any) *BlockQuoteElement {
	self.element.Push(children...)
	return self
}

func (self *BlockQuoteElement) Pop() *BlockQuoteElement {
	self.element.Pop()
	return self
}

func (self BlockQuoteElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self BlockQuoteElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *BlockQuoteElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *BlockQuoteElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
