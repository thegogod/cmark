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
	parent := self.scope
	self.scope = parent.Create()
	nodes := []html.Node{}
	log.Infoln(fmt.Sprintf("entering scope depth %d", self.scope.depth))

	defer func() {
		log.Infoln(fmt.Sprintf("exiting scope depth %d", self.scope.depth))
		self.scope = parent
	}()

	for !scan.Match(RightBrace) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return nil, scan.Curr().Error("expected closing '}'")
		}

		if err != nil {
			return nil, err
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
	for _, statement := range self.statements {
		if err := statement.Validate(scope); err != nil {
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
