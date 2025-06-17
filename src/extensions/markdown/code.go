package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseCode(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseCode(parser, NewScanner(ptr))
}

func (self *Markdown) parseCode(parser html.Parser, scan *_Scanner) (*html.CodeElement, error) {
	code := html.Code()

	if !scan.MatchCount(BackQuote, 1) {
		return code, scan.Curr().Error("expected '`'")
	}

	text, err := self.parseTextUntil(BackQuote, parser, scan)

	if text == nil {
		return code, scan.Curr().Error("expected closing '`'")
	}

	if err != nil {
		return code, err
	}

	log.Debugln("code")
	code.Push(text)
	return code, nil
}
