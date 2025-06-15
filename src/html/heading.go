package html

import (
	"bytes"
	"unicode"

	"github.com/thegogod/cmark/ast"

	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
type HeadingElement struct {
	element *Element
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func Heading(size int, children ...any) *HeadingElement {
	switch size {
	case 1:
		return H1(children...)
	case 2:
		return H2(children...)
	case 3:
		return H3(children...)
	case 4:
		return H4(children...)
	case 5:
		return H5(children...)
	case 6:
		return H6(children...)
	default:
		panic("invalid heading size")
	}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H1(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h1").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H2(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h2").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H3(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h3").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H4(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h4").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H5(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h5").Push(children...)}
}

// https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements
func H6(children ...any) *HeadingElement {
	return &HeadingElement{Elem("h6").Push(children...)}
}

func (self HeadingElement) GetSelector() string {
	return self.element.GetSelector()
}

func (self *HeadingElement) WithAttr(name string, value string) *HeadingElement {
	self.element.WithAttr(name, value)
	return self
}

func (self HeadingElement) HasAttr(name string) bool {
	return self.element.HasAttr(name)
}

func (self HeadingElement) GetAttr(name string) string {
	return self.element.GetAttr(name)
}

func (self *HeadingElement) SetAttr(name string, value string) {
	self.element.SetAttr(name, value)
}

func (self *HeadingElement) DelAttr(name string) {
	self.element.DelAttr(name)
}

func (self *HeadingElement) WithId(value string) *HeadingElement {
	self.element.WithId(value)
	return self
}

func (self HeadingElement) HasId() bool {
	return self.element.HasId()
}

func (self HeadingElement) GetId() string {
	return self.element.GetId()
}

func (self *HeadingElement) SetId(id string) {
	self.element.SetId(id)
}

func (self *HeadingElement) DelId() {
	self.element.DelId()
}

func (self *HeadingElement) WithClass(classes ...string) *HeadingElement {
	self.element.WithClass(classes...)
	return self
}

func (self HeadingElement) HasClass(classes ...string) bool {
	return self.element.HasClass(classes...)
}

func (self HeadingElement) GetClass() []string {
	return self.element.GetClass()
}

func (self *HeadingElement) AddClass(name ...string) {
	self.element.AddClass(name...)
}

func (self *HeadingElement) DelClass(name ...string) {
	self.element.DelClass(name...)
}

func (self *HeadingElement) WithStyles(styles ...maps.KeyValue[string, string]) *HeadingElement {
	self.element.WithStyles(styles...)
	return self
}

func (self HeadingElement) GetStyles() maps.OMap[string, string] {
	return self.element.GetStyles()
}

func (self *HeadingElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	self.element.SetStyles(styles...)
}

func (self *HeadingElement) WithStyle(name string, value string) *HeadingElement {
	self.element.WithStyle(name, value)
	return self
}

func (self HeadingElement) HasStyle(name ...string) bool {
	return self.element.HasStyle(name...)
}

func (self HeadingElement) GetStyle(name string) string {
	return self.element.GetStyle(name)
}

func (self *HeadingElement) SetStyle(name string, value string) {
	self.element.SetStyle(name, value)
}

func (self *HeadingElement) DelStyle(name ...string) {
	self.element.DelStyle(name...)
}

func (self HeadingElement) Count() int {
	return self.element.Count()
}

func (self HeadingElement) Children() []Node {
	return self.element.Children()
}

func (self *HeadingElement) Push(children ...any) *HeadingElement {
	self.element.Push(children...)
	return self
}

func (self *HeadingElement) Pop() *HeadingElement {
	self.element.Pop()
	return self
}

func (self HeadingElement) Render(scope *ast.Scope) []byte {
	if !self.element.attributes.Exists("id") {
		id := []byte{}

		for _, child := range self.element.children {
			value := child.Render(scope)

			for _, b := range value {
				if unicode.IsSpace(rune(b)) {
					id = append(id, '-')
				} else if unicode.IsNumber(rune(b)) {
					id = append(id, b)
				} else if unicode.IsLetter(rune(b)) {
					id = append(id, bytes.ToLower([]byte{b})...)
				}
			}
		}

		self.SetAttr("id", string(id))
	}

	return self.element.Render(scope)
}

func (self HeadingElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	if !self.element.attributes.Exists("id") {
		id := []byte{}

		for _, child := range self.element.children {
			value := child.Render(scope)

			for _, b := range value {
				if unicode.IsSpace(rune(b)) {
					id = append(id, '-')
				} else if unicode.IsNumber(rune(b)) {
					id = append(id, b)
				} else if unicode.IsLetter(rune(b)) {
					id = append(id, bytes.ToLower([]byte{b})...)
				}
			}
		}

		self.SetAttr("id", string(id))
	}

	return self.element.RenderPretty(scope, indent)
}

func (self *HeadingElement) GetById(id string) Node {
	return self.element.GetById(id)
}

func (self *HeadingElement) Select(query ...any) []Node {
	return self.element.Select(query...)
}
