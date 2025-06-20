package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseBlockStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseBlockStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseBlockStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	self.scope = self.scope.Create()
	nodes := []html.Node{}

	defer func() {
		self.scope = self.scope.parent
	}()

	for {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return nil, scan.Curr().Error("expected closing '}'")
		}

		if err != nil {
			return nil, err
		}

		if scan.Curr().Ext() != self.Name() {
			scan.ptr.Back()
			scan.Next()
		}

		if scan.Curr().Kind() == RightBrace {
			break
		}

		nodes = append(nodes, node)
	}

	if _, err := scan.Consume(RightBrace, fmt.Sprintf("expected '}', received '%s'", scan.Curr().String())); err != nil {
		return nil, err
	}

	return BlockStatement{[]Statement{StatementHtml(nodes)}}, nil
}

type BlockStatement struct {
	statements []Statement
}

func (self BlockStatement) Validate(scope *Scope) error {
	child := scope.Create()

	for _, statement := range self.statements {
		if err := statement.Validate(child); err != nil {
			return err
		}
	}

	return nil
}

func (self BlockStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	child := scope.Create()
	values := []string{}

	for _, statement := range self.statements {
		value, err := statement.Evaluate(child)

		if err != nil {
			return value, err
		}

		if !value.IsNil() {
			values = append(values, value.String())
		}
	}

	return reflect.NewString(strings.Join(values, "")), nil
}
