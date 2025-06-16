package markdown

import (
	"slices"

	"github.com/thegogod/cmark/tokens"
	"github.com/thegogod/cmark/tx"
)

type _Scanner struct {
	ptr   *tokens.Pointer
	prev  tokens.Token
	curr  tokens.Token
	types []func(ptr *tokens.Pointer) (*tokens.Token, error)
}

func NewScanner(ptr *tokens.Pointer, types ...func(ptr *tokens.Pointer) (*tokens.Token, error)) *_Scanner {
	return &_Scanner{
		ptr:   ptr,
		types: append(types, tokenScanners...),
	}
}

func (self *_Scanner) Next() bool {
	self.prev = self.curr
	token, err := self.Scan()

	if err != nil {
		return self.Next()
	}

	self.curr = *token
	return token.Kind() > 0
}

func (self *_Scanner) NextWhile(kind ...rune) int {
	i := 0

	for slices.Contains(kind, self.curr.Kind()) {
		i++
		self.Next()
	}

	return i
}

func (self *_Scanner) Match(kind ...rune) bool {
	tx := tx.New(self)

	for _, k := range kind {
		if self.curr.Kind() != k {
			tx.Rollback()
			return false
		}

		self.Next()
	}

	return true
}

func (self *_Scanner) MatchCount(kind rune, count int) bool {
	tx := tx.New(self)

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

func (self *_Scanner) MatchLiteral(value string) bool {
	tx := tx.New(self)
	i := 0

	for i < len(value) {
		for _, b := range self.curr.Bytes() {
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

func (self *_Scanner) Consume(kind rune, message string) (*tokens.Token, error) {
	if self.curr.Kind() == kind {
		self.Next()
		return &self.prev, nil
	}

	return nil, self.curr.Error(message)
}

func (self *_Scanner) Scan() (*tokens.Token, error) {
	self.ptr.Start = self.ptr.End
	self.ptr.Next()
	tx := tx.New(self)

	for _, tokenType := range self.types {
		token, err := tokenType(self.ptr)

		if token != nil || err != nil {
			return token, err
		}

		tx.Rollback()
	}

	return nil, nil
}
