package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/maps"
	"github.com/thegogod/cmark/reflect"
)

type HtmlStatement []Statement

func Html(statements ...Statement) HtmlStatement {
	return statements
}

func (self HtmlStatement) GetSelector() string {
	return ""
}

func (self HtmlStatement) HasAttr(name string) bool {
	return false
}

func (self HtmlStatement) GetAttr(name string) string {
	return ""
}

func (self HtmlStatement) SetAttr(name string, value string) {

}

func (self HtmlStatement) DelAttr(name string) {

}

func (self HtmlStatement) HasId() bool {
	return false
}

func (self HtmlStatement) GetId() string {
	return ""
}

func (self HtmlStatement) SetId(id string) {

}

func (self HtmlStatement) DelId() {

}

func (self HtmlStatement) HasClass(classes ...string) bool {
	return false
}

func (self HtmlStatement) GetClass() []string {
	return []string{}
}

func (self HtmlStatement) AddClass(name ...string) {

}

func (self HtmlStatement) DelClass(name ...string) {

}

func (self HtmlStatement) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self HtmlStatement) SetStyles(styles ...maps.KeyValue[string, string]) {

}

func (self HtmlStatement) HasStyle(name ...string) bool {
	return false
}

func (self HtmlStatement) GetStyle(name string) string {
	return ""
}

func (self HtmlStatement) SetStyle(name string, value string) {

}

func (self HtmlStatement) DelStyle(name ...string) {

}

func (self HtmlStatement) Count() int {
	return len(self)
}

func (self HtmlStatement) Validate(scope *Scope) error {
	for _, statement := range self {
		if err := statement.Validate(scope); err != nil {
			return err
		}
	}

	return nil
}

func (self HtmlStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	child := scope.Create()

	for _, statement := range self {
		value, err := statement.Evaluate(child)

		if err != nil {
			return value, err
		}

		if !value.IsNil() {
			return value, nil
		}
	}

	return reflect.NewNil(), nil
}

func (self HtmlStatement) Render() []byte {
	scope := NewScope()
	value := []byte{}

	for _, statement := range self {
		block, err := statement.Evaluate(scope)

		if err != nil {
			continue
		}

		value = append(value, block.String()...)
	}

	return value
}

func (self HtmlStatement) RenderPretty(indent string) []byte {
	scope := NewScope()
	value := []byte{}

	for _, statement := range self {
		block, err := statement.Evaluate(scope)

		if err != nil {
			continue
		}

		value = append(value, block.String()...)
	}

	return value
}

func (self HtmlStatement) GetById(id string) html.Node {
	return nil
}

func (self HtmlStatement) Select(query ...any) []html.Node {
	return []html.Node{}
}
