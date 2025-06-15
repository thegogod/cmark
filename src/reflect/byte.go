package reflect

func NewByte(value byte) Value {
	return Value{
		_type:  NewByteType(),
		_value: value,
	}
}

func (self Value) ByteType() ByteType {
	return self._type.(ByteType)
}

func (self Value) IsByte() bool {
	return self.Kind() == Byte
}

func (self Value) Byte() byte {
	return self._value.(byte)
}

func (self *Value) SetByte(value byte) {
	self._value = value
}

func (self Value) ByteToString() string {
	return string(self.Byte())
}

func (self Value) ByteToInt() int {
	return int(self.Byte())
}
