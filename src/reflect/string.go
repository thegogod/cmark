package reflect

import "strconv"

func NewString(value string) Value {
	return Value{
		_type:  NewStringType(),
		_value: value,
	}
}

func (self Value) StringType() StringType {
	return self._type.(StringType)
}

func (self Value) IsString() bool {
	return self.Kind() == String
}

func (self Value) String() string {
	return self._value.(string)
}

func (self *Value) SetString(value string) {
	self._value = value
}

func (self Value) SubString(i int, j int) string {
	return self.String()[i:j]
}

func (self Value) Append(value string) {
	self._value = self.String() + value
}

func (self Value) StringToInt() int {
	v, err := strconv.Atoi(self.String())

	if err != nil {
		panic(err)
	}

	return v
}
