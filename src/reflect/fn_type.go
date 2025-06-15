package reflect

import "fmt"

type Param struct {
	Name string
	Type Type
}

type FnType struct {
	name       string
	params     []Param
	returnType Type
}

func NewFnType(name string, params []Param, returnType Type) FnType {
	return FnType{
		name:       name,
		params:     params,
		returnType: returnType,
	}
}

func (self FnType) Kind() Kind {
	return Fn
}

func (self FnType) Name() string {
	return self.name
}

func (self FnType) String() string {
	return fmt.Sprintf("<fn %s>", self.name)
}

func (self FnType) Len() int {
	panic("method not supported")
}

func (self FnType) Comparable() bool {
	return false
}

func (self FnType) Numeric() bool {
	return false
}

func (self FnType) Collection() bool {
	return false
}

func (self FnType) Params() []Param {
	return self.params
}

func (self FnType) ReturnType() Type {
	return self.returnType
}

func (self FnType) Equals(t Type) bool {
	if t.Kind() != Fn {
		return false
	}

	fn := t.(FnType)

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

func (self FnType) ConvertableTo(t Type) bool {
	return false
}

func (self FnType) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]
	return ok
}

func (self FnType) GetMember(name string) Type {
	return memberTypes[self.Kind()][name]
}
