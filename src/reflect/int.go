package reflect

import "strconv"

func NewInt(value int) Value {
	return Value{
		_type:  NewIntType(),
		_value: value,
	}
}

func (self Value) IntType() IntType {
	return self._type.(IntType)
}

func (self Value) IsInt() bool {
	return self.Kind() == Int
}

func (self Value) Int() int {
	return self._value.(int)
}

func (self *Value) SetInt(value int) {
	self._value = value
}

func (self Value) IntToString() string {
	return strconv.Itoa(self.Int())
}
