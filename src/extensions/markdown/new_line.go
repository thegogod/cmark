package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseNewLine(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseNewLine(parser, NewScanner(ptr))
}

func (self *Markdown) parseNewLine(_ html.Parser, scan *_Scanner) (html.Raw, error) {
	if !scan.Match(NewLine) {
		return nil, scan.Curr().Error("expected newline")
	}

	if scan.Match(NewLine) {
		return nil, nil
	}

	for range self.blockQuoteDepth {
		if !scan.Match(GreaterThan) {
			break
		}
	}

	curr := scan.Curr().String()

	if curr == " " || curr == "\n" {
		return nil, nil
	}

	return html.Raw("\n"), nil
}
