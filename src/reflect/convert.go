package reflect

func (self Value) ToString() string {
	if self.IsNil() {
		return self.NilToString()
	}

	if self.IsBool() {
		return self.BoolToString()
	}

	if self.IsByte() {
		return self.ByteToString()
	}

	if self.IsInt() {
		return self.IntToString()
	}

	if self.IsFloat() {
		return self.FloatToString()
	}

	if self.IsFn() {
		return self.FnToString()
	}

	if self.IsNativeFn() {
		return self.NativeFnToString()
	}

	if self.IsMap() {
		return self.MapToString()
	}

	if self.IsMod() {
		return self.ModToString()
	}

	if self.IsSlice() {
		return self.SliceToString()
	}

	return self.String()
}

func (self Value) ToInt() int {
	if self.IsByte() {
		return self.ByteToInt()
	}

	if self.IsFloat() {
		return self.FloatToInt()
	}

	if self.IsString() {
		return self.StringToInt()
	}

	return self.Int()
}
