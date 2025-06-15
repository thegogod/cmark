package reflect

func sliceAt() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("at", []Param{{Name: "i", Type: NewIntType()}}, self.SliceType().Type(), func(args []Value) Value {
			return self.Slice()[args[0].Int()]
		})
	}
}

func sliceLen() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("len", []Param{}, NewIntType(), func(args []Value) Value {
			return NewInt(self.Len())
		})
	}
}
