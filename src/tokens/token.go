package tokens

type Token struct {
	kind  rune
	start Position
	end   Position
	value []byte
}

func New(kind rune, start Position, end Position, value []byte) Token {
	return Token{
		kind:  kind,
		start: start,
		end:   end,
		value: value,
	}
}

func (self Token) Kind() rune {
	return self.kind
}

func (self Token) Start() Position {
	return self.start
}

func (self Token) End() Position {
	return self.end
}

func (self Token) Bytes() []byte {
	return self.value
}

func (self Token) String() string {
	return string(self.value)
}

func (self Token) Error(message string) error {
	return Err(self.start, self.end, message)
}
