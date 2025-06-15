package reflect

type IntType struct{}

func NewIntType() IntType {
	return IntType{}
}

func (self IntType) Kind() Kind {
	return Int
}

func (self IntType) Name() string {
	return Int.String()
}

func (self IntType) String() string {
	return Int.String()
}

func (self IntType) Len() int {
	panic("method not supported")
}

func (self IntType) Comparable() bool {
	return true
}

func (self IntType) Numeric() bool {
	return true
}

func (self IntType) Collection() bool {
	return false
}

func (self IntType) Equals(t Type) bool {
	return t.Kind() == Int
}

func (self IntType) ConvertableTo(t Type) bool {
	kind := t.Kind()
	return kind == Int || kind == Float || kind == String
}

func (self IntType) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]
	return ok
}

func (self IntType) GetMember(name string) Type {
	return memberTypes[self.Kind()][name]
}
