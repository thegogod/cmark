package reflect

type Type interface {
	Kind() Kind
	Name() string
	String() string
	Len() int
	Comparable() bool
	Numeric() bool
	Collection() bool
	Equals(Type) bool
	ConvertableTo(Type) bool
	HasMember(string) bool
	GetMember(string) Type
}

type ComparableType interface {
	Type
	comparable
}

type CallableType interface {
	Type

	Params() []Param
	ReturnType() Type
}
