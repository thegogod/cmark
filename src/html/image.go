package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img
type ImageElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img
func Img() *ImageElement {
	return &ImageElement{Elem("img").Void()}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img#src
func (self *ImageElement) WithSrc(value string) *ImageElement {
	return self.WithAttr("src", value)
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLImageElement/alt#usage_notes
func (self *ImageElement) WithAlt(value string) *ImageElement {
	return self.WithAttr("alt", value)
}

func (self ImageElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *ImageElement) WithAttr(name string, value string) *ImageElement {
	self.element.WithAttr(name, value)
	return self
}

func (self ImageElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self ImageElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *ImageElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *ImageElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *ImageElement) WithId(value string) *ImageElement {
	self.element.WithId(value)
	return self
}

func (self ImageElement) HasId() bool {
	return self.element.HasId()
}

func (self ImageElement) GetId() string {
	return self.element.GetId()
}

func (self *ImageElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *ImageElement) DelId() {
	self.element.DelId()
}

func (self *ImageElement) WithClass(classes ...string) *ImageElement {
	self.element.WithClass(classes...)
	return self
}

func (self ImageElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self ImageElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *ImageElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *ImageElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *ImageElement) WithStyles(styles ...maps.KeyValue[string, string]) *ImageElement {
	self.element.WithStyles(styles...)
	return self
}

func (self ImageElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *ImageElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *ImageElement) WithStyle(name string, value string) *ImageElement {
	self.element.WithStyle(name, value)
	return self
}

func (self ImageElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self ImageElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *ImageElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *ImageElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self ImageElement) Render(scope *ast.Scope) []byte {
	return self.element.Render(scope)
}

func (self ImageElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return self.element.RenderPretty(scope, indent)
}

func (self *ImageElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *ImageElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
