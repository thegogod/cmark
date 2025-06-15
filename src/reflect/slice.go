package reflect

func NewSlice[T Type](_type T, value []Value) Value {
	return Value{
		_type: SliceType{
			_type: _type,
		},
		_value: value,
	}
}

func (self Value) SliceType() SliceType {
	return self._type.(SliceType)
}

func (self Value) IsSlice() bool {
	return self.Kind() == Slice
}

func (self Value) Slice() []Value {
	return self._value.([]Value)
}

func (self Value) At(i int) Value {
	return self.Slice()[i]
}

func (self Value) SubSlice(i int, j int) []Value {
	return self.Slice()[i:j]
}

func (self Value) Push(value Value) {
	if self._type.Equals(value._type) {
		panic("invalid type")
	}

	v := self.Slice()
	v = append(v, value)
	self._value = v
}

func (self Value) Pop() Value {
	v := self.Slice()
	removed := v[len(v)-1]
	v = v[:len(v)-1]
	self._value = v
	return removed
}

func (self Value) SliceToString() string {
	str := "["

	for i, value := range self.Slice() {
		str += value.ToString()

		if i < self.Len()-1 {
			str += ", "
		}
	}

	return str + "]"
}
