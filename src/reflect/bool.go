package reflect

func NewBool(value bool) Value {
	return Value{
		_type:  NewBoolType(),
		_value: value,
	}
}

func (self Value) BoolType() BoolType {
	return self._type.(BoolType)
}

func (self Value) IsBool() bool {
	return self.Kind() == Bool
}

func (self Value) Bool() bool {
	return self._value.(bool)
}

func (self *Value) SetBool(value bool) {
	self._value = value
}

func (self Value) BoolToString() string {
	if self.Bool() {
		return "true"
	}

	return "false"
}
