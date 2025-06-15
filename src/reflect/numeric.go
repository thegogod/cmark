package reflect

func (self Value) Numeric() bool {
	return self._type.Numeric()
}

func (self Value) Add(o Value) Value {
	if self.IsByte() {
		return NewByte(self.Byte() + o.Byte())
	}

	if self.IsInt() {
		return NewInt(self.Int() + o.Int())
	}

	if self.IsFloat() {
		return NewFloat(self.Float() + o.Float())
	}

	panic("method not supported")
}

func (self Value) Subtract(o Value) Value {
	if self.IsByte() {
		return NewByte(self.Byte() - o.Byte())
	}

	if self.IsInt() {
		return NewInt(self.Int() - o.Int())
	}

	if self.IsFloat() {
		return NewFloat(self.Float() - o.Float())
	}

	panic("method not supported")
}

func (self Value) Multiply(o Value) Value {
	if self.IsByte() {
		return NewByte(self.Byte() * o.Byte())
	}

	if self.IsInt() {
		return NewInt(self.Int() * o.Int())
	}

	if self.IsFloat() {
		return NewFloat(self.Float() * o.Float())
	}

	panic("method not supported")
}

func (self Value) Divide(o Value) Value {
	if self.IsByte() {
		return NewByte(self.Byte() / o.Byte())
	}

	if self.IsInt() {
		return NewInt(self.Int() / o.Int())
	}

	if self.IsFloat() {
		return NewFloat(self.Float() / o.Float())
	}

	panic("method not supported")
}

func (self Value) Increment() {
	if self.IsByte() {
		self._value = self.Byte() + 1
		return
	}

	if self.IsInt() {
		self._value = self.Int() + 1
		return
	}

	if self.IsFloat() {
		self._value = self.Float() + 1
		return
	}

	panic("method not supported")
}

func (self Value) Decrement() {
	if self.IsByte() {
		self._value = self.Byte() - 1
		return
	}

	if self.IsInt() {
		self._value = self.Int() - 1
		return
	}

	if self.IsFloat() {
		self._value = self.Float() - 1
		return
	}

	panic("method not supported")
}
