package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type FunctionStatement struct {
	name       tokens.Token
	params     []VariableStatement
	returnType reflect.Type
	body       []Statement
}

func (self FunctionStatement) Validate(scope *Scope) error {
	for _, stmt := range self.body {
		err := stmt.Validate(scope)

		if err != nil {
			return err
		}
	}

	return nil
}

func (self FunctionStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	for _, stmt := range self.body {
		v, err := stmt.Evaluate(scope)

		if err != nil {
			return reflect.NewNil(), err
		}

		if v.IsNil() {
			return v, nil
		}
	}

	return reflect.NewNil(), nil
}

func (self FunctionStatement) Print() {
	self.PrintIndent(0, "  ")
}

func (self FunctionStatement) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[FunctionStatement]: name=\"%s\" return_type=\"%s\"\n", strings.Repeat(indent, depth), self.name.String(), self.returnType.Name())

	for _, param := range self.params {
		param.PrintIndent(depth+1, indent)
	}

	for _, statement := range self.body {
		statement.PrintIndent(depth+1, indent)
	}
}
