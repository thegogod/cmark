package reflect

func NewNativeFn(
	name string,
	params []Param,
	returnType Type,
	cb func([]Value) Value,
) Value {
	return Value{
		_type: NewNativeFnType(
			name,
			params,
			returnType,
		),
		_value: cb,
	}
}

func (self Value) NativeFnType() NativeFnType {
	return self._type.(NativeFnType)
}

func (self Value) IsNativeFn() bool {
	return self.Kind() == NativeFn
}

func (self Value) NativeFn() func([]Value) Value {
	return self._value.(func([]Value) Value)
}

func (self Value) NativeFnToString() string {
	return self.NativeFnType().String()
}
