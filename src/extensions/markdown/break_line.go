package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBreakLine(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBreakLine(parser, NewScanner(ptr))
}

func (self *Markdown) parseBreakLine(_ html.Parser, scan *_Scanner) (*html.BreakLineElement, error) {
	if !scan.MatchCount(Space, 2) && scan.Curr().Kind() == NewLine {
		return nil, scan.Curr().Error("expected two spaces and a newline")
	}

	return html.Br(), nil
}
