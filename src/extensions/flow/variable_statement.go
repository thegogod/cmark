package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseVariableStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseVariableStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseVariableStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	var _type reflect.Type = nil
	var nilable *tokens.Token = nil
	var init Expression = nil

	isSlice := false
	keyword := scan.Prev()
	name, err := scan.Consume(Identifier, "expected variable name")

	if err != nil {
		return nil, err
	}

	if self.scope.HasLocal(name.String()) {
		return nil, keyword.Error("duplicate name")
	}

	if scan.Match(LeftBracket) {
		if _, err = scan.Consume(RightBracket, "expected ']'"); err != nil {
			return nil, err
		}

		isSlice = true
	}

	if scan.Match(Type) || scan.Match(Identifier) {
		kind := scan.Prev()

		if !self.scope.Has(kind.String()) {
			return nil, kind.Error("type '" + kind.String() + "' not found")
		}

		_type = self.scope.Get(kind.String()).Type

		if scan.Match(QuestionMark) {
			nilable = scan.Prev().Ptr()
		}
	}

	if isSlice {
		_type = reflect.NewSliceType(_type, -1)
	}

	if scan.Match(Eq) {
		init, err = self.parseExpression(parser, scan)

		if err != nil {
			return nil, err
		}

		t := init.Type()

		if _type != nil && !_type.Equals(t) {
			return nil, scan.Prev().Error("expected type '" + _type.Name() + "', received '" + t.Name() + "'")
		}

		_type = t
	}

	if _, err = scan.Consume(SemiColon, "expected ';'"); err != nil {
		return nil, err
	}

	self.scope.SetLocal(name.String(), &ScopeEntry{Type: _type})
	log.Infoln(fmt.Sprintf(
		"defined new variable '%s' with type '%s' at depth %d",
		name.String(),
		_type.Name(),
		self.scope.depth,
	))

	return VariableStatement{
		keyword: keyword,
		name:    *name,
		_type:   _type,
		nilable: nilable,
		init:    init,
	}, nil
}

type VariableStatement struct {
	keyword tokens.Token
	name    tokens.Token
	_type   reflect.Type
	nilable *tokens.Token
	init    Expression
}

func (self VariableStatement) Validate(scope *Scope) error {
	if self.init != nil {
		if err := self.init.Validate(scope); err != nil {
			return err
		}
	}

	if scope.HasLocal(self.name.String()) {
		return self.keyword.Error("duplicate name")
	}

	return nil
}

func (self VariableStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	value := reflect.NewNil()

	if self.init != nil {
		v, err := self.init.Evaluate(scope)

		if err != nil {
			return value, err
		}

		value = v
	}

	scope.SetLocal(self.name.String(), &ScopeEntry{Value: value})
	return value, nil
}

func (self VariableStatement) Print() {
	self.PrintIndent(0, "  ")
}

func (self VariableStatement) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[VariableStatement]: keyword=\"%s\" name=\"%s\" type=\"%s\"\n", strings.Repeat(indent, depth), self.keyword.String(), self.name.String(), self._type.Name())

	if self.init != nil {
		self.init.PrintIndent(depth+1, indent)
	}
}
