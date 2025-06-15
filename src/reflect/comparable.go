package reflect

func (self Value) Comparable() bool {
	return self._type.Comparable()
}

func (self Value) Truthy() bool {
	if self.IsNil() {
		return false
	}

	if self.IsBool() {
		return self.Bool()
	}

	if self.IsByte() {
		return self.Byte() > 0
	}

	if self.IsInt() {
		return self.Int() > 0
	}

	if self.IsFloat() {
		return self.Float() > 0
	}

	if self.IsString() {
		return self.Len() > 0
	}

	panic("method not supported")
}

func (self Value) Eq(o Value) bool {
	if self.IsNil() && o.IsNil() {
		return true
	}

	if self.IsBool() {
		return self.Bool() == o.Bool()
	}

	if self.IsByte() {
		return self.Byte() == o.Byte()
	}

	if self.IsInt() {
		return self.Int() == o.Int()
	}

	if self.IsFloat() {
		return self.Float() == o.Float()
	}

	if self.IsString() {
		return self.String() == o.String()
	}

	panic("method not supported")
}

func (self Value) Gt(o Value) bool {
	if self.IsByte() {
		return self.Byte() > o.Byte()
	}

	if self.IsInt() {
		return self.Int() > o.Int()
	}

	if self.IsFloat() {
		return self.Float() > o.Float()
	}

	if self.IsString() {
		return self.String() > o.String()
	}

	panic("method not supported")
}

func (self Value) GtEq(o Value) bool {
	if self.IsByte() {
		return self.Byte() >= o.Byte()
	}

	if self.IsInt() {
		return self.Int() >= o.Int()
	}

	if self.IsFloat() {
		return self.Float() >= o.Float()
	}

	if self.IsString() {
		return self.String() >= o.String()
	}

	panic("method not supported")
}

func (self Value) Lt(o Value) bool {
	if self.IsByte() {
		return self.Byte() < o.Byte()
	}

	if self.IsInt() {
		return self.Int() < o.Int()
	}

	if self.IsFloat() {
		return self.Float() < o.Float()
	}

	if self.IsString() {
		return self.String() < o.String()
	}

	panic("method not supported")
}

func (self Value) LtEq(o Value) bool {
	if self.IsByte() {
		return self.Byte() <= o.Byte()
	}

	if self.IsInt() {
		return self.Int() <= o.Int()
	}

	if self.IsFloat() {
		return self.Float() <= o.Float()
	}

	if self.IsString() {
		return self.String() <= o.String()
	}

	panic("method not supported")
}
