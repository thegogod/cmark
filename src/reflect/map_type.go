package reflect

import "fmt"

type MapType struct {
	key   Type
	value Type
}

func NewMapType(key Type, value Type) *MapType {
	return &MapType{
		key:   key,
		value: value,
	}
}

func (self MapType) Kind() Kind {
	return Map
}

func (self MapType) Name() string {
	return Map.String()
}

func (self MapType) String() string {
	return fmt.Sprintf(
		"%s[%s, %s]",
		Map.String(),
		self.key.Name(),
		self.value.Name(),
	)
}

func (self MapType) Len() int {
	panic("method not supported")
}

func (self MapType) Comparable() bool {
	return false
}

func (self MapType) Numeric() bool {
	return false
}

func (self MapType) Collection() bool {
	return true
}

func (self MapType) Key() Type {
	return self.key
}

func (self MapType) Value() Type {
	return self.value
}

func (self MapType) Equals(t Type) bool {
	if t.Kind() != Map {
		return false
	}

	mp := t.(MapType)

	if !self.key.Equals(mp.key) {
		return false
	}

	if !self.value.Equals(mp.value) {
		return false
	}

	return true
}

func (self MapType) ConvertableTo(t Type) bool {
	return false
}

func (self MapType) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]
	return ok
}

func (self MapType) GetMember(name string) Type {
	return memberTypes[self.Kind()][name]
}
