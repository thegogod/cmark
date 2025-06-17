package assert

import "errors"

type OrStatement []Evaluation

func Or(evaluations ...Evaluation) OrStatement {
	return evaluations
}

func (self OrStatement) Equal(actual any) OrStatement {
	self = append(self, Equal(actual))
	return self
}

func (self OrStatement) Len(length int) OrStatement {
	self = append(self, Len(length))
	return self
}

func (self OrStatement) Nil() OrStatement {
	self = append(self, Nil())
	return self
}

func (self OrStatement) Evaluate(value any) error {
	errs := []error{}

	for _, eval := range self {
		err := eval.Evaluate(value)

		if err == nil {
			return nil
		}

		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
