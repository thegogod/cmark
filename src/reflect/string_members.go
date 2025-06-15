package reflect

import "strings"

func stringAt() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("at", []Param{{Name: "i", Type: NewIntType()}}, NewByteType(), func(args []Value) Value {
			return NewByte(self.String()[args[0].Int()])
		})
	}
}

func stringLen() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("len", []Param{}, NewIntType(), func(args []Value) Value {
			return NewInt(self.Len())
		})
	}
}

func stringReplace() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("replace", []Param{
			{Name: "from", Type: NewStringType()},
			{Name: "to", Type: NewStringType()},
		}, NewStringType(), func(args []Value) Value {
			return NewString(strings.Replace(
				self.String(),
				args[0].String(),
				args[1].String(),
				-1,
			))
		})
	}
}

func stringSlice() MemberCallback {
	return func(self Value) Value {
		return NewNativeFn("slice", []Param{
			{Name: "start", Type: NewIntType()},
			{Name: "end", Type: NewIntType()},
		}, NewStringType(), func(args []Value) Value {
			return NewString(self.String()[args[0].Int():args[1].Int()])
		})
	}
}
