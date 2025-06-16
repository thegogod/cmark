package markdown

import "github.com/thegogod/cmark/tokens"

type _Scanner struct {
	ptr   *tokens.Pointer
	prev  tokens.Token
	curr  tokens.Token
	saves []_Scanner
}

func (self *_Scanner) Next() bool {
	self.prev = self.curr
	token, err := self.Scan()

	if err != nil {
		return self.Next()
	}

	self.curr = token
	return token.Kind() > 0
}

func (self *_Scanner) Scan() (tokens.Token, error) {
	return tokens.Token{}, nil
}
