package reflect

type MemberCallback func(Value) Value

var members = map[Kind]map[string]MemberCallback{
	Bool: {
		"to_string": toString(),
	},
	Byte: {
		"to_string": toString(),
		"to_int":    toInt(),
	},
	Float: {
		"to_string": toString(),
		"to_int":    toInt(),
	},
	Fn: {
		"to_string": toString(),
	},
	Int: {
		"to_string": toString(),
	},
	Map: {
		"to_string": toString(),
	},
	Mod: {
		"to_string": toString(),
	},
	NativeFn: {
		"to_string": toString(),
	},
	Nil: {
		"to_string": toString(),
	},
	Slice: {
		"to_string": toString(),
		"at":        sliceAt(),
		"len":       sliceLen(),
	},
	String: {
		"to_int":  toInt(),
		"at":      stringAt(),
		"len":     stringLen(),
		"replace": stringReplace(),
		"slice":   stringSlice(),
	},
}

var memberTypes = map[Kind]map[string]Type{
	Bool: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Byte: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
		"to_int":    NewNativeFnType("to_int", []Param{}, NewIntType()),
	},
	Float: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
		"to_int":    NewNativeFnType("to_int", []Param{}, NewIntType()),
	},
	Fn: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Int: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Map: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Mod: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	NativeFn: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Nil: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	Slice: {
		"to_string": NewNativeFnType("to_string", []Param{}, NewStringType()),
	},
	String: {
		"to_int": NewNativeFnType("to_int", []Param{}, NewIntType()),
		"at":     NewNativeFnType("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType()),
		"len":    NewNativeFnType("len", []Param{}, NewIntType()),
		"replace": NewNativeFnType(
			"replace",
			[]Param{
				{Name: "from", Type: NewStringType()},
				{Name: "to", Type: NewStringType()},
			},
			NewStringType(),
		),
		"slice": NewNativeFnType(
			"slice",
			[]Param{
				{Name: "start", Type: NewIntType()},
				{Name: "end", Type: NewIntType()},
			},
			NewStringType(),
		),
	},
}

func toString() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("to_string", []Param{}, NewStringType(), func(args []Value) Value {
			return NewString(self.ToString())
		})
	}
}

func toInt() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("to_int", []Param{}, NewIntType(), func(args []Value) Value {
			return NewInt(self.ToInt())
		})
	}
}
