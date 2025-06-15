package reflect

type BoolType struct{}

func NewBoolType() BoolType {
	return BoolType{}
}

func (self BoolType) Kind() Kind {
	return Bool
}

func (self BoolType) Name() string {
	return Bool.String()
}

func (self BoolType) String() string {
	return Bool.String()
}

func (self BoolType) Len() int {
	panic("method not supported")
}

func (self BoolType) Comparable() bool {
	return true
}

func (self BoolType) Numeric() bool {
	return false
}

func (self BoolType) Collection() bool {
	return false
}

func (self BoolType) Equals(t Type) bool {
	return t.Kind() == Bool
}

func (self BoolType) ConvertableTo(t Type) bool {
	kind := t.Kind()
	return kind == Bool || kind == String
}

func (self BoolType) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]
	return ok
}

func (self BoolType) GetMember(name string) Type {
	return memberTypes[self.Kind()][name]
}
