package flow

import (
	"fmt"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParsePrimaryExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parsePrimaryExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parsePrimaryExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	if scan.Match(True) || scan.Match(False) {
		v, err := scan.Prev().Bool()

		if err != nil {
			return nil, scan.Prev().Error(err.Error())
		}

		return LiteralExpression{reflect.NewBool(v)}, nil
	} else if scan.Match(LInt) {
		v, err := scan.Prev().Int()

		if err != nil {
			return nil, scan.Prev().Error(err.Error())
		}

		return LiteralExpression{reflect.NewInt(v)}, nil
	} else if scan.Match(LFloat) {
		v, err := scan.Prev().Float()

		if err != nil {
			return nil, scan.Prev().Error(err.Error())
		}

		return LiteralExpression{reflect.NewFloat(v)}, nil
	} else if scan.Match(LString) {
		return LiteralExpression{reflect.NewString(scan.Prev().String())}, nil
	} else if scan.Match(LByte) {
		return LiteralExpression{reflect.NewByte(scan.Prev().Byte())}, nil
	} else if scan.Match(Nil) {
		return LiteralExpression{reflect.NewNil()}, nil
	} else if scan.Match(Self) {
		if !self.scope.Has("self") {
			return nil, scan.Prev().Error("self is undefined")
		}

		return SelfExpression{
			keyword: scan.Prev(),
			_type:   self.scope.Get("self").Value.(reflect.Type),
		}, nil
	} else if scan.Match(Identifier) {
		if !self.scope.Has(scan.Prev().String()) {
			return nil, scan.Prev().Error("undefined identifier '" + scan.Prev().String() + "'")
		}

		return VariableExpression{
			name:  scan.Prev(),
			_type: self.scope.Get(scan.Prev().String()).Value.(reflect.Type),
		}, nil
	} else if scan.Match(LeftParen) {
		e, err := self.parseExpression(parser, scan)

		if err != nil {
			return nil, err
		}

		scan.Consume(RightParen, "expected ')'")
		return GroupingExpression{e}, nil
	} else if scan.Match(LeftBracket) {
		var _type reflect.Type = nil
		items := []Expression{}

		if !scan.Match(RightBracket) {
			for {
				e, err := self.parseExpression(parser, scan)

				if err != nil {
					return nil, err
				}

				t := e.Type()

				if err = e.Validate(self.scope); err != nil {
					return nil, err
				}

				if _type == nil {
					_type = t
				} else if !_type.Equals(t) {
					return nil, scan.Prev().Error(fmt.Sprintf(
						"expected type '%s', received '%s'",
						_type.Name(),
						t.Name(),
					))
				}

				items = append(items, e)

				if !scan.Match(Comma) {
					break
				}
			}
		}

		if _, err := scan.Consume(RightBracket, "expected ']'"); err != nil {
			return nil, err
		}

		return SliceExpression{
			_type: reflect.NewSliceType(_type, -1),
			items: items,
		}, nil
	}

	return nil, scan.Prev().Error(fmt.Sprintf("expected expression, received '%s' of kind %v", scan.Curr().String(), scan.Curr().Kind()))
}
