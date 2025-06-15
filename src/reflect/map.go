package reflect

import "fmt"

func NewMap[K ComparableType, V Type](key K, value V) Value {
	return Value{
		_type:  NewMapType(key, value),
		_value: map[any]Value{},
	}
}

func (self Value) MapType() MapType {
	return self._type.(MapType)
}

func (self Value) IsMap() bool {
	return self.Kind() == Map
}

func (self Value) Map() map[any]Value {
	return self._value.(map[any]Value)
}

func (self Value) HasKey(key any) bool {
	_, ok := self.Map()[key]
	return ok
}

func (self Value) GetKey(key any) Value {
	return self.Map()[key]
}

func (self *Value) SetKey(key any, value Value) {
	v := self.Map()
	v[key] = value
	self._value = v
}

func (self *Value) DelKey(key any) {
	v := self.Map()
	delete(v, key)
	self._value = v
}

func (self Value) MapToString() string {
	str := "{"
	i := 0

	for key, value := range self.Map() {
		str += fmt.Sprintf("%v: %s", key, value.ToString())

		if i < self.Len()-1 {
			str += ", "
		}

		i++
	}

	return str + "}"
}
