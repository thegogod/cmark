package markdown

import (
	"slices"

	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

type Scanner struct {
	ptr   *tokens.Pointer
	types []func(ptr *tokens.Pointer) (*tokens.Token, error)
}

func NewScanner(ptr *tokens.Pointer) *Scanner {
	self := &Scanner{
		ptr:   ptr,
		types: tokenScanners,
	}

	if ptr.Sof() {
		self.Next()
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
	self.ptr.Start = self.ptr.End
	self.ptr.Next()
	tx := tx.New(self.ptr)

	for _, tokenType := range self.types {
		token, err := tokenType(self.ptr)

		if token != nil || err != nil {
			return token, err
		}

		tx.Rollback()
	}

	return nil, nil
}
