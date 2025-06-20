package flow

import (
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
)

type StatementHtml []html.Node

func (self StatementHtml) Validate(scope *Scope) error {
	return nil
}

func (self StatementHtml) Evaluate(scope *Scope) (reflect.Value, error) {
	values := []string{}

	for _, node := range self {
		values = append(values, string(node.Render()))
	}

	return reflect.NewString(strings.Join(values, "")), nil
}
