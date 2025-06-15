package reflect

type StringType struct{}

func NewStringType() StringType {
	return StringType{}
}

func (self StringType) Kind() Kind {
	return String
}

func (self StringType) Name() string {
	return String.String()
}

func (self StringType) String() string {
	return String.String()
}

func (self StringType) Len() int {
	panic("method not supported")
}

func (self StringType) Comparable() bool {
	return true
}

func (self StringType) Numeric() bool {
	return false
}

func (self StringType) Collection() bool {
	return true
}

func (self StringType) Equals(t Type) bool {
	return t.Kind() == String
}

func (self StringType) ConvertableTo(t Type) bool {
	kind := t.Kind()
	return kind == String || kind == Int || kind == Float || kind == Bool
}

func (self StringType) HasMember(name string) bool {
	_, ok := members[self.Kind()][name]
	return ok
}

func (self StringType) GetMember(name string) Type {
	return memberTypes[self.Kind()][name]
}
