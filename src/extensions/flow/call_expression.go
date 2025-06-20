package flow

import (
	"fmt"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseCallExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseCallExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseCallExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	var _type reflect.Type = nil
	expr, err := self.parsePrimaryExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for {
		if scan.Match(LeftParen) {
			args := []Expression{}

			switch v := expr.(type) {
			case VariableExpression:
				_type = self.scope.Get(v.name.String()).Value.(reflect.Type)
			case GetExpression:
				_type = v.Type()
			}

			fn, ok := _type.(reflect.CallableType)

			if !ok {
				return nil, scan.Prev().Error("expected type 'fn', received '" + _type.Name() + "'")
			}

			if scan.Curr().Kind() != RightParen {
				i := 0

				for {
					arg, err := self.parseExpression(parser, scan)

					if err != nil {
						return nil, err
					}

					args = append(args, arg)

					if i > len(fn.Params())-1 {
						return nil, scan.Prev().Error("too many arguments")
					}

					t := arg.Type()

					if !fn.Params()[i].Type.Equals(t) {
						return nil, scan.Prev().Error(fmt.Sprintf(
							"expected type '%s', received '%s'",
							fn.Params()[i].Type.Name(),
							t.Name(),
						))
					}

					if !scan.Match(Comma) {
						break
					}

					i++
				}
			}

			paren, err := scan.Consume(RightParen, "expected ')'")

			if err != nil {
				return nil, err
			}

			if len(args) != len(fn.Params()) {
				return nil, scan.Prev().Error(fmt.Sprintf(
					"expected %d arguments, received %d",
					len(fn.Params()),
					len(args),
				))
			}

			expr = CallExpression{
				callee: expr,
				paren:  *paren,
				args:   args,
			}

			continue
		} else if scan.Match(Dot) {
			name, err := scan.Consume(Identifier, "expected property name")

			if err != nil {
				return nil, err
			}

			expr = GetExpression{
				object: expr,
				name:   *name,
			}

			continue
		}

		break
	}

	return expr, nil
}

type CallExpression struct {
	callee Expression
	paren  tokens.Token
	args   []Expression
}

func (self CallExpression) Type() reflect.Type {
	t := self.callee.Type()

	if callable, ok := t.(reflect.CallableType); ok {
		t = callable.ReturnType()
	}

	return t
}

func (self CallExpression) Validate(scope *Scope) error {
	t := self.callee.Type()

	if _, ok := t.(reflect.CallableType); !ok {
		return fmt.Errorf("type '%s' is not callable", t.Name())
	}

	return nil
}

func (self CallExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	callee, err := self.callee.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	args := []reflect.Value{}

	for _, arg := range self.args {
		v, err := arg.Evaluate(scope)

		if err != nil {
			return reflect.NewNil(), err
		}

		args = append(args, v)
	}

	if !callee.IsFn() && !callee.IsNativeFn() {
		return reflect.NewNil(), self.paren.Error("expected function")
	}

	if callee.IsNativeFn() {
		return callee.NativeFn()(args), nil
	}

	if len(args) != len(callee.FnType().Params()) {
		return reflect.NewNil(), self.paren.Error(fmt.Sprintf(
			"expected %d arguments, received %d",
			len(callee.FnType().Params()),
			len(args),
		))
	}

	child := scope.Create()

	for i, arg := range args {
		name := callee.FnType().Params()[i].Name
		child.SetLocal(name, &ScopeEntry{
			Name:  name,
			Kind:  VarScope,
			Value: arg,
		})
	}

	fn := callee.Fn().(FunctionStatement)

	if err := fn.Validate(child); err != nil {
		return reflect.NewNil(), err
	}

	return fn.Evaluate(child)
}
