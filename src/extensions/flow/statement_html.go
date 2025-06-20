package flow

import (
	"fmt"
	r "reflect"
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

func (self StatementHtml) Print() {
	self.PrintIndent(0, "  ")
}

func (self StatementHtml) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[StatementHtml]:\n", strings.Repeat(indent, depth))

	for _, node := range self {
		self.printNode(node, depth+1, indent)
	}
}

func (self StatementHtml) printNode(node html.Node, depth int, indent string) {
	content := ""

	if raw, ok := node.(html.Raw); ok {
		content = string(raw)
	}

	fmt.Printf(
		"%s[%s]: %s\n",
		strings.Repeat(indent, depth),
		r.Indirect(r.ValueOf(node)).Type().Name(),
		content,
	)

	if parent, ok := node.(html.ParentNode); ok {
		for _, child := range parent.Children() {
			self.printNode(child, depth+1, indent)
		}
	}
}
