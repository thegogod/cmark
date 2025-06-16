package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseCode(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseCode(parser, NewScanner(ptr))
}

func (self *Markdown) parseCode(parser ast.Parser, scan *_Scanner) (*html.CodeElement, error) {
	if !scan.MatchCount(BackQuote, 1) {
		return nil, scan.Curr().Error("expected '`'")
	}

	code := html.Code()
	text, err := self.parseTextUntil(BackQuote, parser, scan)

	if text == nil {
		return code, scan.Curr().Error("expected closing '`'")
	}

	if err != nil {
		return code, err
	}

	code.Push(text)
	return code, nil
}
