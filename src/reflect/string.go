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
	if self.IsBool() {
		return self.BoolToString()
	} else if self.IsByte() {
		return self.ByteToString()
	} else if self.IsFloat() {
		return self.FloatToString()
	} else if self.IsInt() {
		return self.IntToString()
	} else if self.IsFn() {
		return self.FnToString()
	} else if self.IsMap() {
		return self.MapToString()
	} else if self.IsMod() {
		return self.ModToString()
	} else if self.IsNil() {
		return self.NilToString()
	} else if self.IsSlice() {
		return self.SliceToString()
	}

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
