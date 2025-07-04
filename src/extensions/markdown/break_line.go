package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBreakLine(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBreakLine(parser, NewScanner(ptr))
}

func (self *Markdown) parseBreakLine(_ html.Parser, scan *Scanner) (*html.BreakLineElement, error) {
	el := html.Br()

	if !(scan.MatchCount(Space, 2) && scan.Curr().Kind() == NewLine) {
		return el, scan.Curr().Error("expected two spaces and a newline")
	}

	log.Debugln("break_line")
	return el, nil
}
