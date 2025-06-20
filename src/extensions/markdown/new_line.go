package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseNewLine(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseNewLine(parser, NewScanner(ptr))
}

func (self *Markdown) parseNewLine(_ html.Parser, scan *Scanner) (html.Node, error) {
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

	log.Debugln("new_line")
	return html.Raw("\n"), nil
}
