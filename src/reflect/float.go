package reflect

import "fmt"

func NewFloat(value float64) Value {
	return Value{
		_type:  NewFloatType(),
		_value: value,
	}
}

func (self Value) FloatType() FloatType {
	return self._type.(FloatType)
}

func (self Value) IsFloat() bool {
	return self.Kind() == Float
}

func (self Value) Float() float64 {
	return self._value.(float64)
}

func (self *Value) SetFloat(value float64) {
	self._value = value
}

func (self Value) FloatToString() string {
	return fmt.Sprintf("%f", self.Float())
}

func (self Value) FloatToInt() int {
	return int(self.Float())
}
