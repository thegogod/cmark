package html

import (
	"strings"

	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/reflect"

	"github.com/thegogod/cmark/maps"
)

type Raw []byte

func (self Raw) GetSelector() string {
	return ""
}

func (self Raw) HasAttr(name string) bool {
	return false
}

func (self Raw) GetAttr(name string) string {
	return ""
}

func (self Raw) SetAttr(name string, value string) {
	return
}

func (self Raw) DelAttr(name string) {
	return
}

func (self Raw) HasId() bool {
	return false
}

func (self Raw) GetId() string {
	return ""
}

func (self Raw) SetId(id string) {
	return
}

func (self Raw) DelId() {
	return
}

func (self Raw) HasClass(name ...string) bool {
	return false
}

func (self Raw) GetClass() []string {
	return []string{}
}

func (self Raw) AddClass(name ...string) {
	return
}

func (self Raw) DelClass(name ...string) {
	return
}

func (self Raw) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self Raw) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self Raw) HasStyle(name ...string) bool {
	return false
}

func (self Raw) GetStyle(name string) string {
	return ""
}

func (self Raw) SetStyle(name string, value string) {
	return
}

func (self Raw) DelStyle(name ...string) {
	return
}

func (self Raw) Render(scope *ast.Scope) []byte {
	return self
}

func (self Raw) RenderPretty(scope *ast.Scope, indent string) []byte {
	lines := strings.Split(string(self), "\n")
	return []byte(strings.Join(lines, "\n"+indent))
}

func (self Raw) GetById(id string) Node {
	return nil
}

func (self Raw) Select(query ...any) []Node {
	return []Node{}
}

func (self Raw) Validate(scope *ast.Scope) error {
	return nil
}

func (self Raw) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
