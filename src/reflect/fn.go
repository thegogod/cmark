package reflect

func NewFn(
	name string,
	params []Param,
	returnType Type,
	value any,
) Value {
	return Value{
		_type: NewFnType(
			name,
			params,
			returnType,
		),
		_value: value,
	}
}

func (self Value) FnType() FnType {
	return self._type.(FnType)
}

func (self Value) IsFn() bool {
	return self.Kind() == Fn
}

func (self Value) Fn() any {
	return self._value
}

func (self Value) FnToString() string {
	return self.FnType().String()
}
