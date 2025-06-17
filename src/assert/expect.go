package assert

import (
	"testing"
)

type ExpectStatement struct {
	test        *testing.T
	value       any
	message     string
	evaluations []Evaluation
}

func Expect(value any) *ExpectStatement {
	return &ExpectStatement{value: value}
}

func (self *ExpectStatement) Message(message string) *ExpectStatement {
	self.message = message
	return self
}

func (self *ExpectStatement) And(evaluations ...Evaluation) *ExpectStatement {
	self.evaluations = append(self.evaluations, And(evaluations...))
	return self
}

func (self *ExpectStatement) Or(evaluations ...Evaluation) *ExpectStatement {
	self.evaluations = append(self.evaluations, Or(evaluations...))
	return self
}

func (self *ExpectStatement) Equal(actual any) *ExpectStatement {
	self.evaluations = append(self.evaluations, Equal(actual))
	return self
}

func (self *ExpectStatement) Len(length int) *ExpectStatement {
	self.evaluations = append(self.evaluations, Len(length))
	return self
}

func (self *ExpectStatement) Nil() *ExpectStatement {
	self.evaluations = append(self.evaluations, Nil())
	return self
}

func (self ExpectStatement) Assert(t Test) {
	for _, eval := range self.evaluations {
		err := eval.Evaluate(self.value)

		if err != nil {
			if self.message != "" {
				t.Error(self.message)
			} else {
				t.Error(err)
			}
		}
	}
}

func (self ExpectStatement) AssertNow(t Test) {
	for _, eval := range self.evaluations {
		err := eval.Evaluate(self.value)

		if err != nil {
			if self.message != "" {
				t.Fatal(self.message)
			} else {
				t.Fatal(err)
			}
		}
	}
}
