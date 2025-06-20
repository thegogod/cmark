package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseForStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseForStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseForStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	// self.scope = self.scope.Create()

	// defer func() {
	// 	self.scope = self.scope.parent
	// }()

	var init Statement = nil
	var cond Expression = nil
	var inc Expression = nil

	_, err := scan.Consume(LeftParen, "expected '('")

	if err != nil {
		return nil, err
	}

	if scan.Match(Let) || scan.Match(Const) {
		init, err = self.parseVariableStatement(parser, scan)

		if err != nil {
			return nil, err
		}
	}

	cond, err = self.parseExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	if init != nil {
		if _, err = scan.Consume(SemiColon, "expected ';'"); err != nil {
			return nil, err
		}

		inc, err = self.parseExpression(parser, scan)

		if err != nil {
			return nil, err
		}
	}

	if _, err = scan.Consume(RightParen, "expected ')'"); err != nil {
		return nil, err
	}

	body, err := self.parseStatement(parser, scan)

	if err != nil {
		return nil, err
	}

	if cond == nil {
		cond = LiteralExpression{reflect.NewBool(true)}
	}

	return ForStatement{
		init: init,
		cond: cond,
		inc:  inc,
		body: body,
	}, nil
}

type ForStatement struct {
	init Statement
	cond Expression
	inc  Expression
	body Statement
}

func (self ForStatement) Validate(scope *Scope) error {
	child := scope.Create()

	if self.init != nil {
		if err := self.init.Validate(child); err != nil {
			return err
		}
	}

	if err := self.cond.Validate(child); err != nil {
		return err
	}

	if self.inc != nil {
		if err := self.inc.Validate(child); err != nil {
			return err
		}
	}

	if err := self.body.Validate(child); err != nil {
		return err
	}

	return nil
}

func (self ForStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	child := scope.Create()

	if self.init != nil {
		if _, err := self.init.Evaluate(child); err != nil {
			return reflect.NewString(""), err
		}
	}

	values := []string{}

	for {
		cond, err := self.cond.Evaluate(child)

		if err != nil {
			return reflect.NewString(""), err
		}

		if !cond.Truthy() {
			break
		}

		value, err := self.body.Evaluate(child)

		if err != nil {
			return value, err
		}

		if !value.IsNil() {
			values = append(values, value.String())
		}

		if self.inc != nil {
			if value, err = self.inc.Evaluate(child); err != nil {
				return value, err
			}
		}
	}

	return reflect.NewString(strings.Join(values, "")), nil
}

func (self ForStatement) Print() {
	self.PrintIndent(0, "  ")
}

func (self ForStatement) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[ForStatement]:\n", strings.Repeat(indent, depth))

	if self.init != nil {
		self.init.PrintIndent(depth+1, indent)
	}

	self.cond.PrintIndent(depth+1, indent)

	if self.inc != nil {
		self.inc.PrintIndent(depth+1, indent)
	}

	self.body.PrintIndent(depth+1, indent)
}
