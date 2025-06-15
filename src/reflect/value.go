package reflect

import (
	"fmt"
	"reflect"
)

type Value struct {
	_type  Type
	_value any
}

func (self Value) Type() Type {
	return self._type
}

func (self Value) Kind() Kind {
	return self._type.Kind()
}

func (self Value) Any() any {
	return self._value
}

func (self Value) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]

	if !ok && self.IsMod() {
		return self.HasExport(name)
	}

	return ok
}

func (self Value) GetMember(name string) Value {
	cb, ok := members[self.Kind()][name]

	if !ok && self.IsMod() {
		return self.GetExport(name)
	}

	return cb(self)
}

func ValueOf(value any) Value {
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Bool {
		return NewBool(v.Bool())
	} else if v.CanFloat() {
		return NewFloat(v.Float())
	} else if v.CanInt() {
		return NewInt(int(v.Int()))
	} else if v.Kind() == reflect.String {
		return NewString(v.String())
	} else if (v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer) && v.IsNil() {
		return NewNil()
	}

	panic(fmt.Sprintf("unsupported type: %s", reflect.TypeOf(value).Name()))
}
