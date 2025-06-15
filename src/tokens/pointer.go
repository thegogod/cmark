package tokens

type Pointer struct {
	Src   []byte
	Start Position
	End   Position
}

func Ptr(src []byte) *Pointer {
	return &Pointer{
		Src:   src,
		Start: Position{},
		End:   Position{},
	}
}

func (self Pointer) Sof() bool {
	return self.Start.Index == 0
}

func (self Pointer) Eof() bool {
	return self.End.Index >= len(self.Src)
}

func (self Pointer) Curr() byte {
	if self.Start.Index >= len(self.Src) {
		return 0
	}

	return self.Src[self.Start.Index]
}

func (self Pointer) Peek() byte {
	if self.Eof() {
		return 0
	}

	return self.Src[self.End.Index]
}

func (self *Pointer) Next() byte {
	self.End.Index++
	self.End.Col++

	if self.Peek() == '\n' {
		self.End.Ln++
		self.End.Col = 0
	}

	return self.Peek()
}

func (self Pointer) Bytes() []byte {
	return self.Src[self.Start.Index:self.End.Index]
}

func (self Pointer) Err(message string) error {
	return Err(self.Start, self.End, message)
}

func (self *Pointer) Ok(kind rune) Token {
	return New(
		kind,
		self.Start,
		self.End,
		self.Bytes(),
	)
}
