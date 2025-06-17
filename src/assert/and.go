package assert

type AndStatement []Evaluation

func And(evaluations ...Evaluation) AndStatement {
	return evaluations
}

func (self AndStatement) Equal(actual any) AndStatement {
	self = append(self, Equal(actual))
	return self
}

func (self AndStatement) Len(length int) AndStatement {
	self = append(self, Len(length))
	return self
}

func (self AndStatement) Nil() AndStatement {
	self = append(self, Nil())
	return self
}

func (self AndStatement) Evaluate(value any) error {
	for _, eval := range self {
		err := eval.Evaluate(value)

		if err != nil {
			return err
		}
	}

	return nil
}
