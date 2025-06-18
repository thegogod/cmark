package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH2(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseH2(parser, NewScanner(ptr))
}

func (self *Markdown) parseH2(parser html.Parser, scan *Scanner) (*html.HeadingElement, error) {
	heading := html.H2()

	if !(scan.MatchCount(Hash, 2) && scan.Match(Space)) {
		return heading, scan.Curr().Error("expected '## '")
	}

	for scan.Curr().Kind() != Eof && scan.Curr().Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	log.Debugln("h2")
	return heading, nil
}
