package tokens

import "strconv"

type Token struct {
	kind  rune
	ext   string
	start Position
	end   Position
	value []byte
}

func New(kind rune, ext string, start Position, end Position, value []byte) Token {
	return Token{
		kind:  kind,
		ext:   ext,
		start: start,
		end:   end,
		value: value,
	}
}

func (self Token) Kind() rune {
	return self.kind
}

func (self Token) Ext() string {
	return self.ext
}

func (self Token) Start() Position {
	return self.start
}

func (self Token) End() Position {
	return self.end
}

func (self Token) Error(message string) error {
	return Err(self.start, self.end, message)
}

func (self Token) Ptr() *Token {
	return &self
}

func (self Token) Bytes() []byte {
	return self.value
}

func (self Token) Byte() byte {
	return self.value[0]
}

func (self Token) String() string {
	return string(self.value)
}

func (self Token) Int() (int, error) {
	return strconv.Atoi(string(self.value))
}

func (self Token) Float() (float64, error) {
	return strconv.ParseFloat(string(self.value), 64)
}

func (self Token) Bool() (bool, error) {
	return strconv.ParseBool(string(self.value))
}
