package html

import (
	"fmt"
	"slices"
	"strings"

	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/reflect"

	"github.com/thegogod/cmark/maps"
)

type Element struct {
	Kind       string
	void       bool // https://developer.mozilla.org/en-US/docs/Glossary/Void_element
	attributes maps.OMap[string, string]
	children   []Node
}

func Elem(kind string) *Element {
	return &Element{
		Kind:       kind,
		attributes: maps.OMap[string, string]{},
		children:   []Node{},
	}
}

func (self *Element) Void() *Element {
	self.void = true
	return self
}

func (self Element) GetSelector() string {
	return self.Kind
}

func (self *Element) WithAttr(name string, value string) *Element {
	self.SetAttr(name, value)
	return self
}

func (self Element) HasAttr(name string) bool {
	if self.attributes == nil {
		self.attributes = maps.OMap[string, string]{}
	}

	return self.attributes.Exists(name)
}

func (self Element) GetAttr(name string) string {
	if self.attributes == nil {
		self.attributes = maps.OMap[string, string]{}
	}

	return self.attributes.GetOrDefault(name)
}

func (self *Element) SetAttr(name string, value string) {
	if self.attributes == nil {
		self.attributes = maps.OMap[string, string]{}
	}

	self.attributes.Set(name, value)
}

func (self *Element) DelAttr(name string) {
	if self.attributes == nil {
		self.attributes = maps.OMap[string, string]{}
	}

	self.attributes.Del(name)
}

func (self *Element) WithId(id string) *Element {
	self.SetId(id)
	return self
}

func (self Element) HasId() bool {
	return self.HasAttr("id")
}

func (self Element) GetId() string {
	return self.GetAttr("id")
}

func (self *Element) SetId(id string) {
	self.SetAttr("id", id)
}

func (self *Element) DelId() {
	self.DelAttr("id")
}

func (self *Element) WithClass(name ...string) *Element {
	self.AddClass(name...)
	return self
}

func (self Element) HasClass(name ...string) bool {
	classes := self.GetClass()

	for _, cls := range name {
		if !slices.Contains(classes, cls) {
			return false
		}
	}

	return true
}

func (self Element) GetClass() []string {
	classes := []string{}
	existing := strings.Split(self.GetAttr("class"), " ")

	for _, cls := range existing {
		cls = strings.TrimSpace(cls)

		if len(cls) > 0 {
			classes = append(classes, cls)
		}
	}

	return classes
}

func (self *Element) AddClass(name ...string) {
	classes := self.GetClass()

	for _, cls := range name {
		cls = strings.TrimSpace(cls)

		if len(cls) == 0 {
			continue
		}

		if !slices.Contains(classes, cls) {
			classes = append(classes, cls)
		}
	}

	self.SetAttr("class", strings.Join(classes, " "))
}

func (self *Element) DelClass(name ...string) {
	classes := self.GetClass()
	classes = slices.DeleteFunc(classes, func(cls string) bool {
		return slices.Contains(name, cls)
	})

	self.SetAttr("class", strings.Join(classes, " "))
}

func (self *Element) WithStyles(styles ...maps.KeyValue[string, string]) *Element {
	self.SetStyles(styles...)
	return self
}

func (self Element) GetStyles() maps.OMap[string, string] {
	styles := maps.OMap[string, string]{}
	lines := strings.SplitSeq(self.attributes.GetOrDefault("style"), ";")

	for line := range lines {
		parts := strings.Split(line, ": ")

		if len(parts) != 2 {
			continue
		}

		styles.Set(parts[0], parts[1])
	}

	return styles
}

func (self *Element) SetStyles(styles ...maps.KeyValue[string, string]) {
	value := []string{}

	for _, pair := range styles {
		if pair.Key == "" || pair.Value == "" {
			continue
		}

		value = append(value, fmt.Sprintf("%s: %s", pair.Key, pair.Value))
	}

	if len(value) == 0 {
		self.DelAttr("style")
		return
	}

	self.SetAttr("style", strings.Join(append(value, ""), ";"))
}

func (self *Element) WithStyle(name string, value string) *Element {
	self.SetStyle(name, value)
	return self
}

func (self Element) HasStyle(name ...string) bool {
	styles := self.GetStyles()

	for _, style := range name {
		if !styles.Exists(style) {
			return false
		}
	}

	return true
}

func (self Element) GetStyle(name string) string {
	styles := self.GetStyles()
	return styles.GetOrDefault(name)
}

func (self *Element) SetStyle(name string, value string) {
	styles := self.GetStyles()
	styles.Set(name, value)
	self.SetStyles(styles...)
}

func (self *Element) DelStyle(name ...string) {
	styles := self.GetStyles()

	for _, style := range name {
		styles.Del(style)
	}

	self.SetStyles(styles...)
}

func (self Element) Count() int {
	if self.children == nil {
		self.children = []Node{}
	}

	return len(self.children)
}

func (self Element) Children() []Node {
	if self.children == nil {
		self.children = []Node{}
	}

	return self.children
}

func (self *Element) Push(children ...any) *Element {
	if self.children == nil {
		self.children = []Node{}
	}

	for _, child := range children {
		switch v := child.(type) {
		case Host:
			for key, value := range v {
				self.SetAttr(key, fmt.Sprintf("%v", value))
			}

			break
		case Node:
			self.children = append(self.children, v)
			break
		case string:
			self.children = append(self.children, Raw(v))
			break
		case *string:
			self.children = append(self.children, Raw(*v))
			break
		case []byte:
			self.children = append(self.children, Raw(v))
			break
		default:
			panic("invalid type")
		}
	}

	if len(self.children) > 1 {
		curr, isCurrRaw := self.children[len(self.children)-1].(Raw)
		prev, isPrevRaw := self.children[len(self.children)-2].(Raw)

		if isCurrRaw && isPrevRaw {
			prev = append(prev, curr...)
			self.Pop()
			self.children[len(self.children)-1] = prev
		}
	}

	return self
}

func (self *Element) Pop() *Element {
	if self.children == nil || len(self.children) == 0 {
		return self
	}

	self.children = self.children[:len(self.children)-1]
	return self
}

func (self Element) Render(scope *ast.Scope) []byte {
	html := "<" + self.Kind

	for _, attr := range self.attributes {
		html += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Value)
	}

	if self.void {
		html += " />"
		return []byte(html)
	}

	html += ">"

	for _, child := range self.children {
		if strings.HasPrefix(child.GetSelector(), ":") {
			continue
		}

		html += string(child.Render(scope))
	}

	html += fmt.Sprintf("</%s>", self.Kind)
	return []byte(html)
}

func (self Element) RenderPretty(scope *ast.Scope, indent string) []byte {
	html := "<" + self.Kind

	for _, attr := range self.attributes {
		html += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Value)
	}

	if self.void {
		html += " />"
		return []byte(html)
	}

	html += ">"

	for _, child := range self.children {
		if strings.HasPrefix(child.GetSelector(), ":") {
			continue
		}

		lines := strings.Split(string(child.RenderPretty(scope, indent)), "\n")

		if len(lines) == 3 {
			html += "\n" + indent + string(child.Render(scope))
			continue
		}

		html += "\n" + indent + strings.Join(lines, "\n"+indent)
	}

	html += fmt.Sprintf("\n</%s>", self.Kind)
	return []byte(html)
}

func (self *Element) GetById(id string) Node {
	if self.GetAttr("id") == id {
		return self
	}

	for _, child := range self.children {
		if node := child.GetById(id); node != nil {
			return node
		}
	}

	return nil
}

func (self *Element) Select(query ...any) []Node {
	stmt := Select()

	for _, q := range query {
		switch v := q.(type) {
		case SelectStatement:
			stmt.And(v)
			break
		case string:
			break
		default:
			panic("invalid selector type")
		}
	}

	nodes := []Node{}

	if stmt.Eval(self) {
		nodes = append(nodes, self)
	}

	for _, child := range self.children {
		nodes = append(nodes, child.Select(query...)...)
	}

	return nodes
}

func (self Element) Validate(scope *ast.Scope) error {
	return nil
}

func (self Element) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
