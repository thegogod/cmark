package html

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

type MetaDataElement map[string]any

func MetaData() MetaDataElement {
	return MetaDataElement{}
}

func (self MetaDataElement) Exists(key string) bool {
	_, exists := self[key]
	return exists
}

func (self MetaDataElement) Get(key string) any {
	value, exists := self[key]

	if exists {
		return value
	}

	return nil
}

func (self MetaDataElement) Set(key string, value any) MetaDataElement {
	self[key] = value
	return self
}

func (self MetaDataElement) GetSelector() string {
	return ":metadata"
}

func (self MetaDataElement) HasAttr(name string) bool {
	return false
}

func (self MetaDataElement) GetAttr(name string) string {
	return ""
}

func (self MetaDataElement) SetAttr(name string, value string) {
	return
}

func (self MetaDataElement) DelAttr(name string) {
	return
}

func (self MetaDataElement) HasId() bool {
	return false
}

func (self MetaDataElement) GetId() string {
	return ""
}

func (self MetaDataElement) SetId(id string) {
	return
}

func (self MetaDataElement) DelId() {
	return
}

func (self MetaDataElement) HasClass(name ...string) bool {
	return false
}

func (self MetaDataElement) GetClass() []string {
	return []string{}
}

func (self MetaDataElement) AddClass(name ...string) {
	return
}

func (self MetaDataElement) DelClass(name ...string) {
	return
}

func (self MetaDataElement) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self MetaDataElement) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self MetaDataElement) HasStyle(name ...string) bool {
	return false
}

func (self MetaDataElement) GetStyle(name string) string {
	return ""
}

func (self MetaDataElement) SetStyle(name string, value string) {
	return
}

func (self MetaDataElement) DelStyle(name ...string) {
	return
}

func (self MetaDataElement) Render(scope *ast.Scope) []byte {
	return []byte{}
}

func (self MetaDataElement) RenderPretty(scope *ast.Scope, indent string) []byte {
	return []byte{}
}

func (self MetaDataElement) GetById(id string) Node {
	return nil
}

func (self MetaDataElement) Select(query ...any) []Node {
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

	return nodes
}

func (self MetaDataElement) Validate(scope *ast.Scope) error {
	return nil
}

func (self MetaDataElement) Evaluate(scope *ast.Scope) (reflect.Value, error) {
	value := self.Render(scope)
	return reflect.NewString(string(value)), nil
}
