package reflect

import "fmt"

type NativeFnType struct {
	name       string
	params     []Param
	returnType Type
}

func NewNativeFnType(name string, params []Param, returnType Type) NativeFnType {
	return NativeFnType{
		name:       name,
		params:     params,
		returnType: returnType,
	}
}

func (self NativeFnType) Kind() Kind {
	return NativeFn
}

func (self NativeFnType) Name() string {
	return self.name
}

func (self NativeFnType) String() string {
	return fmt.Sprintf("<native_fn %s>", self.name)
}

func (self NativeFnType) Len() int {
	panic("method not supported")
}

func (self NativeFnType) Comparable() bool {
	return false
}

func (self NativeFnType) Numeric() bool {
	return false
}

func (self NativeFnType) Collection() bool {
	return false
}

func (self NativeFnType) Params() []Param {
	return self.params
}

func (self NativeFnType) ReturnType() Type {
	return self.returnType
}

func (self NativeFnType) Equals(t Type) bool {
	if t.Kind() != NativeFn {
		return false
	}

	fn := t.(NativeFnType)

	if len(self.params) != len(fn.params) {
		return false
	}

	if !self.returnType.Equals(fn.returnType) {
		return false
	}

	for i, param := range self.params {
		if !param.Type.Equals(fn.params[i].Type) {
			return false
		}
	}

	return true
}

func (self NativeFnType) ConvertableTo(t Type) bool {
	return false
}

func (self NativeFnType) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]
	return ok
}

func (self NativeFnType) GetMember(name string) Type {
	return memberTypes[self.Kind()][name]
}
