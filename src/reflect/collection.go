package reflect

func (self Value) Collection() bool {
	return self._type.Collection()
}

func (self Value) Len() int {
	if self.IsMap() {
		return len(self._value.(map[any]Value))
	}

	if self.IsMod() {
		return len(self._value.(map[string]Value))
	}

	if self.IsSlice() {
		return len(self.Slice())
	}

	if self.IsString() {
		return len(self.String())
	}

	panic("method not supported")
}
