package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseNewLine(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseNewLine(parser, NewScanner(ptr))
}

func (self *Markdown) parseNewLine(_ html.Parser, scan *_Scanner) (html.Raw, error) {
	el := html.Raw("\n")

	if !scan.Match(NewLine) {
		return el, scan.Curr().Error("expected newline")
	}

	for range self.blockQuoteDepth {
		if !scan.Match(GreaterThan) {
			break
		}
	}

	log.Debugln("new_line")
	return el, nil
}
