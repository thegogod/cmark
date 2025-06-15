package html

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/ast"

	"github.com/thegogod/cmark/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Guides/Comments
type Comment []byte

func (self Comment) GetSelector() string {
	return ""
}

func (self Comment) HasAttr(name string) bool {
	return false
}

func (self Comment) GetAttr(name string) string {
	return ""
}

func (self Comment) SetAttr(name string, value string) {
	return
}

func (self Comment) DelAttr(name string) {
	return
}

func (self Comment) HasId() bool {
	return false
}

func (self Comment) GetId() string {
	return ""
}

func (self Comment) SetId(id string) {
	return
}

func (self Comment) DelId() {
	return
}

func (self Comment) HasClass(name ...string) bool {
	return false
}

func (self Comment) GetClass() []string {
	return []string{}
}

func (self Comment) AddClass(name ...string) {
	return
}

func (self Comment) DelClass(name ...string) {
	return
}

func (self Comment) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self Comment) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self Comment) HasStyle(name ...string) bool {
	return false
}

func (self Comment) GetStyle(name string) string {
	return ""
}

func (self Comment) SetStyle(name string, value string) {
	return
}

func (self Comment) DelStyle(name ...string) {
	return
}

func (self Comment) Render(scope *ast.Scope) []byte {
	return []byte(fmt.Sprintf("<!-- %s -->", string(self)))
}

func (self Comment) RenderPretty(scope *ast.Scope, indent string) []byte {
	lines := strings.Split(string(self), "\n")
	formatted := strings.Join(lines, "\n"+indent)
	return []byte(fmt.Sprintf("<!--\n %s \n-->", formatted))
}

func (self Comment) GetById(id string) Node {
	return nil
}

func (self Comment) Select(query ...any) []Node {
	return []Node{}
}
