package flow

import (
	"slices"

	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

type Scanner struct {
	ptr *tokens.Pointer
}

func NewScanner(ptr *tokens.Pointer) *Scanner {
	self := &Scanner{ptr}

	if ptr.Sof() {
		self.Scan()
	}

	return self
}

func (self Scanner) Ptr() *tokens.Pointer {
	return self.ptr
}

func (self Scanner) Prev() tokens.Token {
	return self.ptr.Iter.Prev
}

func (self Scanner) Curr() tokens.Token {
	return self.ptr.Iter.Curr
}

func (self *Scanner) Next() bool {
	token, err := self.Scan()

	if err != nil {
		return self.Next()
	}

	return token.Kind() > 0
}

func (self *Scanner) NextWhile(kind ...rune) int {
	i := 0

	for slices.Contains(kind, self.Curr().Kind()) {
		i++
		self.Next()
	}

	return i
}

func (self *Scanner) Match(kind ...rune) bool {
	tx := tx.New(self.ptr)

	for _, k := range kind {
		if self.Curr().Kind() != k {
			tx.Rollback()
			return false
		}

		self.Next()
	}

	return true
}

func (self *Scanner) MatchCount(kind rune, count int) bool {
	tx := tx.New(self.ptr)

	for range count {
		if !self.Match(kind) {
			tx.Rollback()
			return false
		}
	}

	if self.Match(kind) {
		tx.Rollback()
		return false
	}

	return true
}

func (self *Scanner) MatchLiteral(value string) bool {
	tx := tx.New(self.ptr)
	i := 0

	for i < len(value) {
		for _, b := range self.Curr().Bytes() {
			if i >= len(value) {
				break
			}

			if b != value[i] {
				tx.Rollback()
				return false
			}

			i++
		}

		if !self.Next() {
			tx.Rollback()
			return false
		}
	}

	return true
}

func (self *Scanner) Consume(kind rune, message string) (*tokens.Token, error) {
	if self.Curr().Kind() == kind {
		self.Next()
		return self.Prev().Ptr(), nil
	}

	return nil, self.Curr().Error(message)
}

func (self *Scanner) Scan() (*tokens.Token, error) {
	if self.ptr.Eof() {
		return self.ptr.Ok(Eof).Ptr(), nil
	}

	self.ptr.Start = self.ptr.End
	b := self.ptr.Curr()
	self.ptr.Next()

	switch b {
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		// ignore whitespace
		break
	case '@':
		return self.ptr.Ok(At).Ptr(), nil
	case '(':
		return self.ptr.Ok(LeftParen).Ptr(), nil
	case ')':
		return self.ptr.Ok(RightParen).Ptr(), nil
	case '{':
		return self.ptr.Ok(LeftBrace).Ptr(), nil
	case '}':
		return self.ptr.Ok(RightBrace).Ptr(), nil
	case '[':
		return self.ptr.Ok(LeftBracket).Ptr(), nil
	case ']':
		return self.ptr.Ok(RightBracket).Ptr(), nil
	case ',':
		return self.ptr.Ok(Comma).Ptr(), nil
	case '.':
		return self.ptr.Ok(Dot).Ptr(), nil
	case ':':
		if self.ptr.Peek() == ':' {
			self.ptr.Next()
			return self.ptr.Ok(DoubleColon).Ptr(), nil
		}

		return self.ptr.Ok(Colon).Ptr(), nil
	case ';':
		return self.ptr.Ok(SemiColon).Ptr(), nil
	case '?':
		return self.ptr.Ok(QuestionMark).Ptr(), nil
	case '|':
		if self.ptr.Peek() != '|' {
			return nil, self.ptr.Err("expected '|'")
		}

		self.ptr.Next()
		return self.ptr.Ok(Or).Ptr(), nil
	case '&':
		if self.ptr.Peek() != '&' {
			return nil, self.ptr.Err("expected '&'")
		}

		self.ptr.Next()
		return self.ptr.Ok(And).Ptr(), nil
	case '+':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(PlusEq).Ptr(), nil
		}

		return self.ptr.Ok(Plus).Ptr(), nil
	case '-':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(MinusEq).Ptr(), nil
		} else if self.ptr.Peek() == '>' {
			self.ptr.Next()
			return self.ptr.Ok(ReturnType).Ptr(), nil
		} else if self.isInt(self.ptr.Peek()) {
			self.ptr.Next()
			return self.onNumeric()
		}

		return self.ptr.Ok(Minus).Ptr(), nil
	case '*':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(StarEq).Ptr(), nil
		}

		return self.ptr.Ok(Star).Ptr(), nil
	case '/':
		if self.ptr.Peek() == '/' {
			return self.onComment()
		} else if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(SlashEq).Ptr(), nil
		}

		return self.ptr.Ok(Slash).Ptr(), nil
	case '!':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(NotEq).Ptr(), nil
		}

		return self.ptr.Ok(Not).Ptr(), nil
	case '=':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(EqEq).Ptr(), nil
		}

		return self.ptr.Ok(Eq).Ptr(), nil
	case '>':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(GtEq).Ptr(), nil
		}

		return self.ptr.Ok(Gt).Ptr(), nil
	case '<':
		if self.ptr.Peek() == '=' {
			self.ptr.Next()
			return self.ptr.Ok(LtEq).Ptr(), nil
		}

		return self.ptr.Ok(Lt).Ptr(), nil
	case '\'':
		return self.onByte()
	case '"':
		return self.onString()
	default:
		if self.isInt(b) {
			return self.onNumeric()
		} else if self.isAlpha(b) {
			return self.onIdentifier()
		}

		return nil, self.ptr.Err("unexpected character")
	}

	return self.Scan()
}

func (self Scanner) isInt(b byte) bool {
	return b >= '0' && b <= '9'
}

func (self Scanner) isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b == '_' || b == '$')
}

func (self *Scanner) onComment() (*tokens.Token, error) {
	for self.ptr.Peek() != '\n' {
		if self.ptr.Eof() {
			return nil, self.ptr.Err("unterminated comment")
		}

		self.ptr.Next()
	}

	self.ptr.Next()
	return self.Scan()
}

func (self *Scanner) onByte() (*tokens.Token, error) {
	self.ptr.Start = self.ptr.End
	self.ptr.Next()

	if self.ptr.Peek() != '\'' {
		return nil, self.ptr.Err("unterminated byte")
	}

	token := self.ptr.Ok(LByte).Ptr()
	self.ptr.Next()
	return token, nil
}

func (self *Scanner) onString() (*tokens.Token, error) {
	self.ptr.Start = self.ptr.End

	for self.ptr.Peek() != '"' {
		if self.ptr.Eof() {
			return nil, self.ptr.Err("unterminated string")
		}

		self.ptr.Next()
	}

	token := self.ptr.Ok(LString).Ptr()
	self.ptr.Next()
	return token, nil
}

func (self *Scanner) onNumeric() (*tokens.Token, error) {
	kind := LInt

	for self.isInt(self.ptr.Peek()) {
		self.ptr.Next()
	}

	if self.ptr.Peek() == '.' {
		kind = LFloat
		self.ptr.Next()

		for self.isInt(self.ptr.Peek()) {
			self.ptr.Next()
		}
	}

	return self.ptr.Ok(kind).Ptr(), nil
}

func (self *Scanner) onIdentifier() (*tokens.Token, error) {
	for self.isAlpha(self.ptr.Peek()) || self.isInt(self.ptr.Peek()) {
		self.ptr.Next()
	}

	name := self.ptr.Bytes()

	if kind, ok := Keywords[string(name)]; ok {
		return self.ptr.Ok(kind).Ptr(), nil
	}

	return self.ptr.Ok(Identifier).Ptr(), nil
}
